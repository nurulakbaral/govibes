package user

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type ResponseRegister struct {
	Id        uuid.UUID   `json:"id"`
	Name      string      `json:"name"`
	Username  string      `json:"username"`
	Email     string      `json:"email"`
	CreatedAt time.Time   `json:"created_at"`
	DeletedAt pgtype.Date `json:"deleted_at"`
}

type RequestRegister struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseLogin = ResponseRegister
