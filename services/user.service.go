package services

import "github.com/ChristinaKozi/go-api/models"

// * is a pointer to an object or type and is being taken in as a parameter
// object after invoking parenthesis defines the type of the return object
// second paranthesis holds returns if returning multiple objects
type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}
