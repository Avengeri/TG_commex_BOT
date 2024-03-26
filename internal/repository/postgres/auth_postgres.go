package postgres

import (
	"TG_commex_BOT/internal/model"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) SetUser(user *model.User) error {
	sqlStatement := `INSERT INTO t_users_auth(id, username, password_hash, chatId)  VALUES ($1,$2,$3,$4)`
	_, err := r.db.Exec(sqlStatement, user.Id, user.UserName, user.PasswordHash, user.ChatId)
	if err != nil {
		return err
	}
	return nil

}

func (r *AuthPostgres) CheckUser(user *model.User) (bool, error) {
	sqlStatement := `SELECT COUNT(*) FROM t_users_auth WHERE username=$1`

	var count int
	err := r.db.QueryRow(sqlStatement, user.UserName).Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (r *AuthPostgres) GetUserInfo(user *model.User) (*model.User, error) {
	var userCopy *model.User
	sqlStatement := `SELECT id FROM t_users_auth WHERE username=$1`
	err := r.db.Get(&userCopy, sqlStatement, user.UserName)
	if err != nil {
		return nil, err
	}
	return userCopy, nil
}

func (r *AuthPostgres) SetUserPassword(password string) error {
	sqlStatement := `INSERT INTO t_users_auth(password_hash) VALUES($1)`
	_, err := r.db.Exec(sqlStatement, password)
	if err != nil {
		return err
	}
	return nil
}
