package main

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func uploadHandler(fsvr *fileServer, w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	files := r.MultipartForm.File["file"]
	var wg sync.WaitGroup
	httpStatus := make(chan int, len(files))

	for _, f := range files {
		wg.Add(1)
		go uploadFiles(f, fsvr.directory, &wg, httpStatus)
	}

	wg.Wait()
	close(httpStatus)
	for code := range httpStatus {
		if code != http.StatusOK {
			w.WriteHeader(code)
			return
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func uploadFiles(h *multipart.FileHeader, dir string, wg *sync.WaitGroup, status chan int) {
	defer wg.Done()

	dstF, err := os.CreateTemp(dir, createTempFilename(h.Filename))
	if err != nil {
		status <- http.StatusBadRequest
		log.Println(err)
		return
	}

	defer ioCloser(dstF)
	srcF, err := h.Open()
	if err != nil {
		status <- http.StatusInternalServerError
		log.Println(err)
		return
	}

	defer ioCloser(srcF)
	if _, err := io.Copy(dstF, srcF); err != nil {
		status <- http.StatusInternalServerError
		log.Println(err)
		return
	}

	status <- http.StatusOK
	log.Printf("file uploaded %s %s\n", h.Filename, humanReadSize(h.Size))
}

func createTempFilename(name string) string {
	ext := filepath.Ext(name)
	return fmt.Sprintf("%s_tmp-*%s", strings.TrimSuffix(name, ext), ext)
}

func ioCloser(i io.Closer) {
	if err := i.Close(); err != nil {
		log.Println(err)
	}
}
