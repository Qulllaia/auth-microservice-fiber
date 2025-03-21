package dto

type UserGetDto struct {
	ID    int    `db:"id"`
	Email string `db:"email" json:"email"`
}
