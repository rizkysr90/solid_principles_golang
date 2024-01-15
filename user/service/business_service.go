package service

import (
	"solid_principles/user"
	"solid_principles/user/repository"
)

// BusinessService orchestrates business operations involving users.
type BusinessService struct {
	UserRepo   repository.UserRepository
	NotifyServ NotifyService
}

// CreateUser saves a user to the database and sends a welcome notification.
func (bs BusinessService) CreateUser(user user.User) {
	bs.UserRepo.SaveUser(user)
	bs.NotifyServ.SendNotification(user, "Welcome to our platform!")
}
