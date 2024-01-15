package service

import (
	"fmt"
	"solid_principles/user"
)

// NotifyService is responsible for sending notifications.
type NotifyService struct{}

// SendNotification sends a notification to the user.
func (n NotifyService) SendNotification(user user.User, message string) {
	// Notification logic...
	fmt.Printf("Notification sent to user %s: %s\n", user.Name, message)
}
