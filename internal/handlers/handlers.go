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

func IndexHandle(res http.ResponseWriter, req *http.Request) {
	req.Header.Add("Content-Type", "text/html")
	http.ServeFile(res, req, "../index.html")
}

func UploadHandle(res http.ResponseWriter, req *http.Request) {

	if err := req.ParseMultipartForm(10 << 20); err != nil {
		http.Error(res, "parse file error", http.StatusInternalServerError)
		return
	}

	file, handler, err := req.FormFile("myFile")
	if err != nil {
		http.Error(res, "file get error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(res, "file read error", http.StatusInternalServerError)
		return
	}

	convertedData := service.Conv(string(data))

	ext := filepath.Ext(handler.Filename)
	filename := fmt.Sprintf("convFile_%s%s", time.Now().UTC().Format("2006-01-02 15-04-05"), ext)

	localFile, err := os.Create(filename)
	if err != nil {
		http.Error(res, "file create error", http.StatusInternalServerError)
		return
	}
	defer localFile.Close()

	if _, err := localFile.WriteString(convertedData); err != nil {
		http.Error(res, "file write error", http.StatusInternalServerError)
		return
	}
	contentType := "text/plain"
	res.Header().Set("Content-Type", contentType)

	res.WriteHeader(http.StatusOK)
	if _, err := res.Write([]byte(convertedData)); err != nil {
		fmt.Println("response write error:", err)
	}
}
