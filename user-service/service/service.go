package service

import (
	"user/model"
	"user/repository"
)

type UserService interface {
	GetAllUsers() ([]*model.User, error)
	GetUserByID(id int64) (*model.User, error)
	GetUserByUserName(username string) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id int64) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (u *userService) GetAllUsers() ([]*model.User, error) {
	return u.repo.GetAllUsers()
}

func (u *userService) GetUserByID(id int64) (*model.User, error) {
	return u.repo.GetUserByID(id)
}

func (u *userService) GetUserByUserName(username string) (*model.User, error) {
	return u.repo.GetUserByUserName(username)
}

func (u *userService) CreateUser(user *model.User) error {
	return u.repo.CreateUser(user)
}

func (u *userService) UpdateUser(user *model.User) error {
	return u.repo.UpdateUser(user)
}

func (u *userService) DeleteUser(id int64) error {
	return u.repo.DeleteUser(id)
}
