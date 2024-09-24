package services

import (
	"github.com/ChristinaKozi/go-api/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            contxt.Context
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	return nil
}

func (u *UserServiceImpl) GetUser(name *string) (*models.User, error) {
	return nil, nil
}

func (u *UserServiceImpl) GetAll() []*models.User {
	return nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	return nil
}

func (u *UserServiceImpl) DeleteUser(name *string) error {
	return nil
}
