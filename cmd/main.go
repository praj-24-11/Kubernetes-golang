package main

import (
	"go-upload-api/internal"
	"log"
	"net/http"
	"os"
)

func main() {
	uploadPath := os.Getenv("UPLOAD_PATH")
	if uploadPath == "" {
		uploadPath = "/data/uploads"
	}

	http.HandleFunc("/upload", internal.UploadHandler(uploadPath))
	http.HandleFunc("/files", internal.ListFilesHandler(uploadPath))
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	log.Println("✅ Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("❌ Failed to start: %v", err)
	}
}
