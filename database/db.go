package database

import (
	"fmt"
	"main/database/dto"
	"main/models"

	"github.com/jmoiron/sqlx"
)

type Database struct {
	db *sqlx.DB
}

func InitDatabse(db *sqlx.DB) *Database{
	return &Database{db:db}
}

func (db *Database) CreateUser(user *models.UserData) error{
	_, err := db.db.Exec(`INSERT INTO "User" (login, password) VALUES ($1, $2)`, user.Login, user.Password);
	return err; 
}

func (db *Database) GetUserWithId(params int) (*dto.UserGetDto, error){
	var user dto.UserGetDto
	err := db.db.Get(&user, `SELECT user_id, login FROM "User" WHERE user_id = $1`, params);
	return &user, err; 
}

func (db *Database) GetUsers() (*[]dto.UserGetDto, error){
	users := []dto.UserGetDto{}
	err := db.db.Select(&users, `SELECT user_id, login FROM "User"`);
	fmt.Println(users[0].Login);
	return &users, err; 
}

func (db *Database) PutUser(userParams models.UserData) error{
	_, err := db.db.Exec(`UPDATE "User" SET login = $1, password = $2 WHERE user_id = $3`, userParams.Login, userParams.Password, userParams.User_id);
	return err; 
}

func (db *Database) DeleteUser(param int) error{
	_, err := db.db.Exec(`DELETE FROM "User" WHERE user_id = $1`, param);
	return err; 
}

func (db *Database) LoginUser(loginParams dto.UserLoginDto) (bool, error){

	var result bool = true;

	err := db.db.Get(&loginParams, `SELECT login, password FROM "User" WHERE login=$1 AND password=$2`, loginParams.Login, loginParams.Password)

	if (dto.UserLoginDto{}) == loginParams{
		result = false
	}

	return result, err
}