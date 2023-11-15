package useCases

import (
	"encoding/json"
	"fmt"
	"go-http/database"
	"go-http/handler"
	"go-http/models"
	"go-http/router"
	"net/http"
)

var GetUsersUseCase = router.Route{
	Path:    `/users`,
	Method:  http.MethodGet,
	Handler: getUsersHandler,
}

func getUsersHandler(w http.ResponseWriter, r *http.Request, quit chan<- bool) {
	users, err := database.GetEntitiesFromDb[[]models.User]()

	if err != nil {
		fmt.Println("GetUsersUseCase", err.Error())
		handler.ErrorHandler(w, err.Error(), http.StatusInternalServerError)
	} else {
		encodedBytes, _ := json.Marshal(users)
		w.Write(encodedBytes)
	}

	quit <- true
}
