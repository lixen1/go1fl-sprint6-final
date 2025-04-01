package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func MainHanler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	file, err := os.ReadFile("../index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(file)
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/upload" {
		http.NotFound(w, r)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Ошибка парсинга формы", http.StatusInternalServerError)
		return
	}

	downloadedFile, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "ошибка при получении файла", http.StatusInternalServerError)
		return
	}
	defer downloadedFile.Close()

	data, err := io.ReadAll(downloadedFile)
	if err != nil {
		http.Error(w, "ошибка конвертации файла в []byte", http.StatusInternalServerError)
		return
	}

	text := service.DataConvert(string(data))

	filename := time.Now().Format("02-01-2006 15:04:05") + filepath.Ext(header.Filename)

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0775)
	if err != nil {
		http.Error(w, "ошибка создания файла результирующего файла", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		http.Error(w, "ошибка записи строки в результирующий файл", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	if _, err := w.Write([]byte(text)); err != nil {
		log.Printf("ошибка отправки запроса %v", err)
	}
}
