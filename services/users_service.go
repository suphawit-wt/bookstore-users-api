package services

import (
	models "bookstore/users/models/users"
	"bookstore/users/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	GetById(int64) (*models.User, *errors.RestErr)
	Search(string) ([]models.UserPublic, *errors.RestErr)
	Create(models.User) *errors.RestErr
	Update(int64, models.User) *errors.RestErr
	Delete(int64) *errors.RestErr
}

func (s *usersService) GetById(userId int64) (*models.User, *errors.RestErr) {
	user := models.User{
		Id: userId,
	}
	if err := user.GetById(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *usersService) Search(status string) ([]models.UserPublic, *errors.RestErr) {
	user := models.User{}
	return user.GetByStatus(status)
}

func (s *usersService) Create(user models.User) *errors.RestErr {
	if err := user.Create(); err != nil {
		return err
	}

	return nil
}

func (s *usersService) Update(userId int64, user models.User) *errors.RestErr {
	user.Id = userId
	if err := user.GetById(); err != nil {
		return err
	}
	if err := user.Update(); err != nil {
		return err
	}

	return nil
}

func (s *usersService) Delete(userId int64) *errors.RestErr {
	user := models.User{
		Id: userId,
	}
	if err := user.GetById(); err != nil {
		return err
	}
	if err := user.Delete(); err != nil {
		return err
	}

	return nil
}
