package useCases

import (
	"go-http/database"
	"go-http/handler"
	"go-http/models"
	"go-http/router"
	"go-http/utilities"
	"net/http"
)

var EditUserUseCase = router.Route{
	Path:    "/users",
	Method:  http.MethodPut,
	Handler: editUserHandler,
}

func editUserHandler(w http.ResponseWriter, r *http.Request, quit chan<- bool) {
	id, getIdErr := utilities.GetIdFromPath(r)
	if getIdErr != nil {
		handler.ErrorHandler(w, getIdErr.Error(), http.StatusBadRequest)
		quit <- true
		return
	}

	userInput, getBodyErr := utilities.ReadBody[models.User](r)
	userInput.Id = id
	if getBodyErr != nil {
		handler.ErrorHandler(w, getBodyErr.Error(), http.StatusInternalServerError)
	}

	_, updateErr := database.Db.Model(&userInput).Column("first_name", "last_name", "age").WherePK().Update()
	if updateErr != nil {
		handler.ErrorHandler(w, updateErr.Error(), http.StatusInternalServerError)
		quit <- true
		return
	}

	quit <- true
}
