package route

import (
	"backenedprojectphase1/go-backend-daily-tasks/day01-create-user-api/handler"
	"net/http"
)

func RegisterUserRoutes(userHandler *handler.UserHandler) {
	http.HandleFunc("/users", userHandler.CreateUser)
	http.HandleFunc("/users/{id}", userHandler.GetUserById)
	http.HandleFunc("/users/{id}", userHandler.DeleteUserBYId)
}
