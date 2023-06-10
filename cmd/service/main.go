package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/progressive-download/", func(w http.ResponseWriter, r *http.Request) {
		filePath := strings.TrimPrefix(r.URL.Path, "/media/")
		if filePath == "" {
			http.Error(w, "File path missing", http.StatusNotFound)
			return
		}
		log.Println("Progressively downloading", filePath)

		file, err := os.Open("media/" + filePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		fileInfo, err := file.Stat()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "video/mp4")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))
		w.Header().Set("Accept-Ranges", "bytes")

		http.ServeContent(w, r, fileInfo.Name(), fileInfo.ModTime(), file)
	})

	http.ListenAndServe(":8080", nil)
}
