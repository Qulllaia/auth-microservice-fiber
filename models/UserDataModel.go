package models

type UserData struct {
	User_id  int    `db:"user_id"`
	Login    string `db:"login" json:"login"`
	Password string `db:"password" json:"password"`
}
