package services

import (
	"bookstore/users/logger"
	"bookstore/users/repository"
	"bookstore/users/utils"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return userService{userRepo: userRepo}
}

func (s userService) GetUser(id int64) (*repository.User, error) {
	user, err := s.userRepo.GetById(id)
	if err != nil {
		logger.Error("Database Error when trying to retrive user by id.", err)
		return nil, utils.SQLServerErrorValidate(err)
	}

	return user, nil
}

func (s userService) SearchUsers(status string) ([]repository.UserPublic, error) {
	users, err := s.userRepo.GetByStatus(status)
	if err != nil {
		logger.Error("Database Error when trying to retrive users by status.", err)
		return nil, utils.SQLServerErrorValidate(err)
	}

	return users, nil
}

func (s userService) CreateUser(req repository.User) error {
	err := s.userRepo.Create(req)
	if err != nil {
		logger.Error("Database Error when trying to insert new user.", err)
		return utils.SQLServerErrorValidate(err)
	}

	return nil
}

func (s userService) UpdateUser(id int64, req repository.User) error {
	_, err := s.userRepo.GetById(id)
	if err != nil {
		logger.Error("Database Error when trying to retrive user by id.", err)
		return utils.SQLServerErrorValidate(err)
	}

	err = s.userRepo.Update(id, req)
	if err != nil {
		logger.Error("Database Error when trying to update user.", err)
		return utils.SQLServerErrorValidate(err)
	}

	return nil
}

func (s userService) DeleteUser(id int64) error {
	_, err := s.userRepo.GetById(id)
	if err != nil {
		logger.Error("Database Error when trying to retrive user by id.", err)
		return utils.SQLServerErrorValidate(err)
	}

	err = s.userRepo.Delete(id)
	if err != nil {
		logger.Error("Database Error when trying to delete user.", err)
		return utils.SQLServerErrorValidate(err)
	}

	return nil
}
