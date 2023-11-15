package useCases

import (
	"encoding/json"
	"go-http/database"
	"go-http/handler"
	"go-http/router"
	"go-http/utilities"
	"net/http"
)

var GetUserUseCase = router.Route{
	Path:    `/users/\d+`,
	Method:  http.MethodGet,
	Handler: getUserHandler,
}

func getUserHandler(w http.ResponseWriter, r *http.Request, quit chan<- bool) {
	id, getIdErr := utilities.GetIdFromPath(r)
	if getIdErr != nil {
		handler.ErrorHandler(w, getIdErr.Error(), http.StatusBadRequest)
		quit <- true
		return
	}

	user, getUserErr := database.GetUserById(id)
	if getUserErr != nil {
		handler.ErrorHandler(w, getUserErr.Error(), http.StatusInternalServerError)
		quit <- true
		return
	}

	encodedBytes, _ := json.Marshal(user)
	w.Write(encodedBytes)
	quit <- true
}
