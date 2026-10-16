package main

import (
	"backenedprojectphase1/go-backend-daily-tasks/day01-create-user-api/config"
	"backenedprojectphase1/go-backend-daily-tasks/day01-create-user-api/handler"
	"backenedprojectphase1/go-backend-daily-tasks/day01-create-user-api/repository"
	"backenedprojectphase1/go-backend-daily-tasks/day01-create-user-api/route"
	"backenedprojectphase1/go-backend-daily-tasks/day01-create-user-api/service"
	"fmt"
	"log"
	"net/http"
)

func main() {

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	userrepo := repository.NewUserRepository(db)
	userservice := service.NewUserService(userrepo)
	userhandler := handler.NewUserHandler(userservice)

	route.RegisterUserRoutes(userhandler)

	fmt.Println("Server started running on port 8080")
	http.ListenAndServe(":8080", nil)

}
