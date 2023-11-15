package useCases

import (
	"fmt"
	"go-http/handler"
	"go-http/router"
	"net/http"
	"os"
	"strings"
)

var DownloadFileUseCase = router.Route{
	Path:    `/file/(\w+)`,
	Method:  http.MethodGet,
	Handler: downloadFileHandler,
}

func downloadFileHandler(w http.ResponseWriter, r *http.Request, quit chan<- bool) {
	fullPath := strings.Split(r.URL.Path, "/")
	filename := fullPath[len(fullPath)-1]
	fileBytes, err := os.ReadFile(fmt.Sprintf("./uploads/%v", filename))

	if err != nil {
		fmt.Println(err.Error())
		handler.ErrorHandler(w, fmt.Sprintf("Could not get resource %v", filename), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%v\"", filename))
	w.Write(fileBytes)
	quit <- true
}
