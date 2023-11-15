package useCases

import (
	"fmt"
	"go-http/handler"
	"go-http/router"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var UploadFileUseCase = router.Route{
	Path:    `/file`,
	Method:  "POST",
	Handler: uploadFileHandler,
}

func uploadFileHandler(w http.ResponseWriter, r *http.Request, quit chan<- bool) {
	parseMultiformError := r.ParseMultipartForm(3000000) // allow 3MB files

	if parseMultiformError != nil {
		handler.ErrorHandler(w, fmt.Sprintf("Could not parse form '%v'", parseMultiformError.Error()), http.StatusInternalServerError)
		return
	}

	file, fileHeader, formFileError := r.FormFile("file")

	if formFileError != nil {
		handler.ErrorHandler(w, formFileError.Error(), http.StatusInternalServerError)
		return
	}

	mkdirError := os.MkdirAll("./uploads", os.ModePerm)
	if mkdirError != nil {
		handler.ErrorHandler(w, mkdirError.Error(), http.StatusInternalServerError)
		return
	}

	nameOfSavedFile := fmt.Sprintf("%d%s", time.Now().Nanosecond(), filepath.Ext(fileHeader.Filename))
	destination, createDestinationError := os.Create(fmt.Sprintf("./uploads/%v", nameOfSavedFile))
	if createDestinationError != nil {
		handler.ErrorHandler(w, createDestinationError.Error(), http.StatusInternalServerError)
		return
	}

	_, copyError := io.Copy(destination, file)
	if copyError != nil {
		handler.ErrorHandler(w, copyError.Error(), http.StatusInternalServerError)
		return
	}

	destination.Close()
	file.Close()
	w.Write([]byte(nameOfSavedFile))
	quit <- true
}
