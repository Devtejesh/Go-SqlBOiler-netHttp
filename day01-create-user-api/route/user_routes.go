package route

import (
	"backenedprojectphase1/go-backend-daily-tasks/day01-create-user-api/handler"
	"net/http"
)

func RegisterUserRoutes(userHandler *handler.UserHandler) {
	http.HandleFunc("POST /users", userHandler.CreateUser)
	http.HandleFunc("GET /users/{id}", userHandler.GetUserById)
	http.HandleFunc("DELETE /users/{id}", userHandler.DeleteUserBYId)
	http.HandleFunc("GET /users", userHandler.GetAllUsers)

}
