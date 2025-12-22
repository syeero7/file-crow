package main

import (
	"io"
	"log"
	"net/http"
)

func streamHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	ft, exist := transfers.get(id)
	if !exist {
		http.Error(w, "transfer not found", http.StatusNotFound)
		return
	}

	if _, err := io.Copy(ft.session.writer, r.Body); err != nil {
		log.Printf("transfer failed: %v", err)
	}

	ft.session.writer.Close()
	<-ft.session.done
	w.WriteHeader(http.StatusOK)
}
