package useCases

import (
	"go-http/database"
	"go-http/handler"
	"go-http/models"
	"go-http/router"
	"go-http/utilities"
	"net/http"
	"strconv"
)

var CreateUserUseCase = router.Route{
	Path:    `/users`,
	Method:  http.MethodPost,
	Handler: createUserHandler,
}

func createUserHandler(w http.ResponseWriter, r *http.Request, quit chan<- bool) {
	user, readBodyErr := utilities.ReadBody[models.User](r)
	if readBodyErr != nil {
		handler.ErrorHandler(w, readBodyErr.Error(), http.StatusInternalServerError)
		quit <- true
		return
	}

	_, insertError := database.Db.Model(&user).Insert()
	if insertError != nil {
		handler.ErrorHandler(w, insertError.Error(), http.StatusInternalServerError)
		quit <- true
		return
	}

	stringifiesId := strconv.Itoa(user.Id)
	idAsBytes := []byte(stringifiesId)
	w.WriteHeader(http.StatusCreated)
	w.Write(idAsBytes)
	quit <- true
}
