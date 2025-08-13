package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Used not allowed method", http.StatusMethodNotAllowed)
		return
	}
	data, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write(data)
}
func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Used not allowed method", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "not enought memory", http.StatusInternalServerError)
		return
	}
	if r.MultipartForm != nil {
		defer r.MultipartForm.RemoveAll()
	}
	file, header, err := r.FormFile("file")
	if err != nil || file == nil {
		if r.MultipartForm != nil {
			for _, fhs := range r.MultipartForm.File {
				if len(fhs) > 0 {
					header = fhs[0]
					file, err = header.Open()
					break
				}
			}
		}

	}
	if err != nil || file == nil {
		http.Error(w, "failed to get file", http.StatusInternalServerError)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "can not read file", http.StatusInternalServerError)
		return
	}

	result, _, err := service.Convert(string(data))
	if err != nil {
		http.Error(w, "failed to Convert", http.StatusInternalServerError)
		return
	}
	ext := filepath.Ext(header.Filename)
	ts := time.Now().UTC().String()
	fname := ts + ext
	out, err := os.Create(fname)
	if err != nil {
		http.Error(w, "failed to create result file", http.StatusInternalServerError)
		return
	}
	defer out.Close()
	if _, err := out.Write([]byte(result)); err != nil {
		http.Error(w, "failed to write result file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = w.Write([]byte(result))
}
