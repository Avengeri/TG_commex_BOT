package repository

import (
	"TG_commex_BOT/internal/model"
	"TG_commex_BOT/internal/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type AuthUser interface {
	SetUser(user *model.User) error
	SetUserPassword(password string) error
	CheckUser(user *model.User) (bool, error)
	GetUserInfo(user *model.User) (*model.User, error)
}

type UserRepository struct {
	Auth AuthUser
}

func NewStorageUserPostgres(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		Auth: postgres.NewAuthPostgres(db),
	}
}
