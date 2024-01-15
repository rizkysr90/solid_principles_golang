package main

import (
	"solid_principles/user"
	"solid_principles/user/repository"
	"solid_principles/user/service"
)

func main() {
	userService := service.BusinessService{
		UserRepo:   repository.UserRepository{},
		NotifyServ: service.NotifyService{},
	}

	newUser := user.User{ID: 1, Name: "John Doe"}
	userService.CreateUser(newUser)
}
