package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type ProgressWriter struct {
	writer     io.Writer
	total      int
	written    int
	limiter    rate.Limiter
	onProgress func(int, int)
}

func (pw *ProgressWriter) Write(d []byte) (int, error) {
	n, err := pw.writer.Write(d)
	if err != nil {
		return n, err
	}

	pw.written += n
	if pw.limiter.Allow() || pw.written == pw.total {
		pw.onProgress(pw.written, pw.total)
	}

	return n, nil
}

type FileProgress struct {
	ID      string `json:"id"`
	Total   int    `json:"total"`
	Current int    `json:"current"`
	Type    string `json:"type"`
}

func streamHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	ft, exist := transfers.get(id)
	if !exist {
		http.Error(w, "transfer not found", http.StatusNotFound)
		return
	}

	ctx := r.Context()
	go func() {
		<-ctx.Done()
		ft.session.writer.CloseWithError(errors.New("disconnected"))
	}()

	pw := newProgressWriter(id, int(r.ContentLength), ft.session.writer)
	buf1MB := make([]byte, 1024*1024)
	if _, err := io.CopyBuffer(pw, r.Body, buf1MB); err != nil {
		if !errors.Is(err, io.ErrClosedPipe) {
			log.Printf("transfer failed: %v", err)
		}
		ft.session.writer.CloseWithError(err)
	}

	if pw.total == pw.written {
		ft.session.writer.Close()
		w.WriteHeader(http.StatusOK)
	}

	<-ft.session.done
}

func newProgressWriter(id string, t int, w io.Writer) *ProgressWriter {
	return &ProgressWriter{
		total:   t,
		writer:  w,
		limiter: *rate.NewLimiter(rate.Every(250*time.Millisecond), 1),
		onProgress: func(current, total int) {
			p := FileProgress{ID: id, Total: total, Current: current, Type: "progress"}
			msg, err := json.Marshal(p)
			if err != nil {
				log.Println(err)
				return
			}
			fileServer.broadcast(msg)
		},
	}
}
