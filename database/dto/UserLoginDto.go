package dto

type UserLoginDto struct {
	Login    string `db:"login" json:"login"`
	Password string `db:"password" json:"password"`
}
