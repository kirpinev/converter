package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

const folderName = "strings"

func GetPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, filepath.Join("..", "index.html"))
}

func GetFormData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "error while parsing form", http.StatusInternalServerError)
		return
	}

	formFile, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "error while trying to find field", http.StatusBadRequest)
		return
	}
	defer formFile.Close()

	filename := handler.Filename
	extension := filepath.Ext(filename)
	if extension == "" {
		extension = ".txt"
	}

	fileBytes, err := io.ReadAll(formFile)

	if err != nil {
		http.Error(w, "error while reading file", http.StatusInternalServerError)
		return
	}

	fileData := service.ConvertString(string(fileBytes))
	ts := time.Now().UTC().Format("20060102T150405Z0700")
	filePath := filepath.Join(folderName, fmt.Sprint(ts, extension))

	err = os.MkdirAll(folderName, 0755)
	if err != nil {
		http.Error(w, "error while creating dir", http.StatusInternalServerError)
		return
	}

	err = os.WriteFile(filePath, []byte(fileData), 0644)
	if err != nil {
		http.Error(w, "error while writing to file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fileData))
}
