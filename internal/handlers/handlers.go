package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

type Handlers struct {
	log *log.Logger
}

func New(logger *log.Logger) *Handlers {
	return &Handlers{log: logger}
}

func (h *Handlers) MainPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func (h *Handlers) UploadFile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		h.log.Println("ошибка парсинга формы:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		h.log.Println("ошибка получения файла из формы:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		h.log.Println("ошибка чтения файла:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := service.DefineOutput(string(data))
	if err != nil {
		h.log.Println("ошибка конвертации:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(header.Filename)
	fileName := fmt.Sprintf("%s%s", time.Now().UTC().String(), ext)

	outFile, err := os.Create(fileName)
	if err != nil {
		h.log.Println("ошибка создания файла:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	if _, err := outFile.WriteString(result); err != nil {
		h.log.Println("ошибка записи в файл:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	if _, err := w.Write([]byte(result)); err != nil {
		h.log.Println("ошибка записи ответа:", err)
	}
}
