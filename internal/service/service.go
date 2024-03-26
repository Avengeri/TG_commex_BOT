package service

import (
	"TG_commex_BOT/internal/model"
	"TG_commex_BOT/internal/repository"
)

type AuthUserService interface {
	SetUserService(user *model.User) error
	SetUserPassword(password string) error
	CheckUserService(user *model.User) (bool, error)
	GetUserInfoService(user *model.User) (*model.User, error)
}

type UserService struct {
	Auth AuthUserService
}

func NewUserService(repos *repository.UserRepository) *UserService {
	return &UserService{
		Auth: NewAuthService(repos.Auth),
	}
}
