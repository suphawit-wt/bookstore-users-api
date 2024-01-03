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

func (s userService) GetUser(id int64) (*UserResponse, error) {
	user, err := s.userRepo.GetById(id)
	if err != nil {
		logger.Error("Database Error when trying to retrive user by id.", err)
		return nil, utils.SQLServerErrorValidate(err)
	}

	userResponse := UserResponse{
		Id:          user.Id,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Status:      user.Status,
		DateCreated: user.DateCreated,
	}

	return &userResponse, nil
}

func (s userService) SearchUsers(status string) ([]SearchUserResponse, error) {
	users, err := s.userRepo.GetByStatus(status)
	if err != nil {
		logger.Error("Database Error when trying to retrive users by status.", err)
		return nil, utils.SQLServerErrorValidate(err)
	}

	searchUserResponseList := []SearchUserResponse{}
	for _, user := range users {
		searchUserResponse := SearchUserResponse{
			Id:          user.Id,
			Status:      user.Status,
			DateCreated: user.DateCreated,
		}

		searchUserResponseList = append(searchUserResponseList, searchUserResponse)
	}

	return searchUserResponseList, nil
}

func (s userService) CreateUser(req CreateUserRequest) error {
	user := repository.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}

	err := s.userRepo.Create(user)
	if err != nil {
		logger.Error("Database Error when trying to insert new user.", err)
		return utils.SQLServerErrorValidate(err)
	}

	return nil
}

func (s userService) UpdateUser(id int64, req UpdateUserRequest) error {
	_, err := s.userRepo.GetById(id)
	if err != nil {
		logger.Error("Database Error when trying to retrive user by id.", err)
		return utils.SQLServerErrorValidate(err)
	}

	user := repository.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}

	err = s.userRepo.Update(id, user)
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
