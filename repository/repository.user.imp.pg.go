package repository

import (
	"database/sql"
	"github.com/alunogabrielgava123/APRENDENDO-GO-2022/model"
)

type ImpUserRepositoryPg struct {
	db *sql.DB
}

func (u *ImpUserRepositoryPg) CreateUser(newUser model.User) (model.User, error) {

	var userResponse model.User

	query := u.db.QueryRow("INSERT INTO album (title, artist, price) VALUES ($1, $2, $3)", newUser.UserName, newUser.Pws)
	if err := query.Scan(&userResponse.ID, &userResponse.UserName); err != nil {
		return userResponse, err
	}

	return userResponse, nil
}

func (u *ImpUserRepositoryPg) DeleteUser(id int64) error {

	_, err := u.db.Exec("DELETE FROM TABLE user WHERE id_user =$1", id)
	if err != nil {
		return err
	}

	return nil

}
