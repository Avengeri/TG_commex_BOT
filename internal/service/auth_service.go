package service

import (
	"TG_commex_BOT/internal/model"
	"TG_commex_BOT/internal/repository"
	"errors"
)

type AuthService struct {
	repo repository.AuthUser
}

func NewAuthService(repo repository.AuthUser) *AuthService {
	return &AuthService{repo: repo}
}

func (r *AuthService) SetUserService(user *model.User) error {
	err := r.repo.SetUser(user)
	if err != nil {
		return errors.New("error when adding a user")
	}
	return nil
}

func (r *AuthService) CheckUserService(user *model.User) (bool, error) {
	exist, err := r.repo.CheckUser(user)
	if err != nil {
		return false, errors.New("error checking the user")
	}
	if !exist {
		return false, errors.New("the user has been not found")
	}
	return true, nil
}

func (r *AuthService) GetUserInfoService(user *model.User) (*model.User, error) {
	user, err := r.repo.GetUserInfo(user)
	if err != nil {
		return nil, errors.New("error getting user")
	}
	return user, nil
}
func (r *AuthService) SetUserPassword(password string) error {
	err := r.repo.SetUserPassword(password)
	if err != nil {
		return errors.New("error add password")
	}
	return nil
}
