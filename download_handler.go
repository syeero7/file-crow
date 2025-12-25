package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type TransferState struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	ft, exist := transfers.get(id)
	if !exist {
		http.Error(w, "transfer not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", ft.name))
	w.Header().Set("Content-Type", "application/octet-stream")
	if _, err := io.Copy(w, ft.session.reader); err != nil {
		log.Printf("transfer failed: %v", err)
		return
	}

	// TODO: handle tcp write: broken pipe error

	ft.session.reader.Close()
	ft.session.done <- struct{}{}
	transfers.remove(id)

	ts := TransferState{ID: id, Type: "done"}
	msg, err := json.Marshal(ts)
	if err != nil {
		log.Printf("failed to encode json: %v", err)
		return
	}
	fileServer.broadcast(msg)
}
