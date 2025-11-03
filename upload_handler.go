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

	for _, f := range files {
		wg.Add(1)
		go func(h *multipart.FileHeader) {
			defer wg.Done()

			dstF, err := os.CreateTemp(fsvr.directory, createTempFilename(h.Filename))
			if err != nil {
				log.Println(err)
				return
			}

			defer ioCloser(dstF)
			srcF, err := h.Open()
			if err != nil {
				log.Println(err)
				return
			}

			defer ioCloser(srcF)
			if _, err := io.Copy(dstF, srcF); err != nil {
				log.Println(err)
				return
			}

			log.Printf("file uploaded %s %s\n", h.Filename, humanReadSize(h.Size))
		}(f)
	}

	wg.Wait()
	http.Redirect(w, r, "/", http.StatusSeeOther)
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
