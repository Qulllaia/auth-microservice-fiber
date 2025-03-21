package database

import (
	"main/database/dto"
	"main/models"
	"os"

	"github.com/jmoiron/sqlx"
)

type Database struct {
	db *sqlx.DB
}

func InitDatabse(db *sqlx.DB) *Database{
	return &Database{db:db}
}

func (db *Database) CreateUser(user *models.UserData) error{
	encryptedPassword, err := Encrypt([]byte(os.Getenv("SECRET_PASSWORD")), user.Password);
	if err != nil{
		return err;
	}
	_, err = db.db.Exec(`INSERT INTO "user" (email, password) VALUES ($1, $2)`, user.Email, encryptedPassword);
	return err; 
}

func (db *Database) GetUserWithId(params int) (*dto.UserGetDto, error){
	var user dto.UserGetDto
	err := db.db.Get(&user, `SELECT id, email FROM "user" WHERE id = $1`, params);
	return &user, err; 
}

func (db *Database) GetUsers() (*[]dto.UserGetDto, error){
	users := []dto.UserGetDto{}
	err := db.db.Select(&users, `SELECT id, email FROM "user"`);
	return &users, err; 
}

func (db *Database) PutUser(userParams models.UserData) error{
	_, err := db.db.Exec(`UPDATE "user" SET email = $1, password = $2 WHERE id = $3`, userParams.Email, userParams.Password, userParams.ID);
	return err; 
}

func (db *Database) DeleteUser(param int) error{
	_, err := db.db.Exec(`DELETE FROM "user" WHERE id = $1`, param);
	return err; 
}

func (db *Database) LoginUser(EmailParams dto.UserLoginDto) (bool, error){

	var result bool = true;

	err := db.db.Get(&EmailParams, `SELECT email, password FROM "user" WHERE Email=$1 AND password=$2`, EmailParams.Email, EmailParams.Password)

	if (dto.UserLoginDto{}) == EmailParams{
		result = false
	}

	return result, err
}