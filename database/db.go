package database

import (
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

func (db *Database) GetUser(params int) (*models.UserData, error){
	var user models.UserData
	err := db.db.Get(&user, `SELECT user_id, login FROM "User" WHERE user_id = $1`, params);
	return &user, err; 
}
