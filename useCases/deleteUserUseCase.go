package useCases

import (
	"go-http/database"
	"go-http/handler"
	"go-http/models"
	"go-http/router"
	"go-http/utilities"
	"net/http"
)

var DeleteUserUseCase = router.Route{
	Path:    `/users/\d+`,
	Method:  http.MethodDelete,
	Handler: deleteUserHandler,
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request, quit chan<- bool) {
	id, getIdErr := utilities.GetIdFromPath(r)
	if getIdErr != nil {
		handler.ErrorHandler(w, getIdErr.Error(), http.StatusBadRequest)
		quit <- true
		return
	}

	_, deleteErr := database.Db.Model(&models.User{Id: id}).WherePK().Delete()

	if deleteErr != nil {
		handler.ErrorHandler(w, deleteErr.Error(), http.StatusBadRequest)
	}
	quit <- true
}
