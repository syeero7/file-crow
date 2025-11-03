package main

import (
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func downloadHandler(fsvr *fileServer, w http.ResponseWriter, r *http.Request) {
	filename := r.PathValue("file")
	if strings.ContainsRune(filename, filepath.Separator) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file := path.Join(fsvr.directory, filename)
	if _, err := os.Stat(file); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, file)
}
