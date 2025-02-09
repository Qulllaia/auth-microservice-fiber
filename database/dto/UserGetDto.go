package dto

type UserGetDto struct {
	User_id int    `db:"user_id"`
	Login   string `db:"login" json:"login"`
}
