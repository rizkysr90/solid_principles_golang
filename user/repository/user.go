package repository

import (
	"fmt"
	"solid_principles/user"
)

// UserRepository is responsible for database interactions related to users.
type UserRepository struct{}

// SaveUser saves a user to the database.
func (r UserRepository) SaveUser(user user.User) {
	// Database logic to save user...
	fmt.Printf("User %s saved to the database.\n", user.Name)
}
