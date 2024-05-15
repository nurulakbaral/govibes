package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type ResponseUserRegister struct {
	Id        uuid.UUID   `json:"id"`
	Name      string      `json:"name"`
	Username  string      `json:"username"`
	Email     string      `json:"email"`
	CreatedAt time.Time   `json:"created_at"`
	DeletedAt pgtype.Date `json:"deleted_at"`
}

type RequestUserRegister struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
