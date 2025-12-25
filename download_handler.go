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

	broadcastTransferState(id, "ready")

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", ft.name))
	w.Header().Set("Content-Type", "application/octet-stream")
	if _, err := io.Copy(w, ft.session.reader); err != nil {
		log.Printf("transfer failed: %v", err)
		broadcastTransferState(id, "failed")
	} else {
		defer broadcastTransferState(id, "done")
	}

	ft.session.reader.Close()
	ft.session.done <- struct{}{}
	transfers.remove(id)
}

func broadcastTransferState(id, message string) {
	ts := TransferState{ID: id, Type: message}
	msg, err := json.Marshal(ts)
	if err != nil {
		log.Printf("failed to encode json: %v", err)
		return
	}
	fileServer.broadcast(msg)
}
