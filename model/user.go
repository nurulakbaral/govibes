package model

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"govibes.app/database"
)

type RequestUser struct {
	Name     string
	Username string
	Email    string
	Password string
}

type ResponseUser struct {
	Id        string      `json:"id"`
	Name      string      `json:"name"`
	Username  string      `json:"username"`
	Email     string      `json:"email"`
	CreatedAt time.Time   `json:"created_at"`
	DeletedAt pgtype.Date `json:"deleted_at"`
}

type User struct {
	Id        string      `json:"id"`
	Name      string      `json:"name"`
	Username  string      `json:"username"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
	CreatedAt time.Time   `json:"created_at"`
	DeletedAt pgtype.Date `json:"deleted_at"`
}

func (user User) InsertUser(ctx context.Context) error {
	userArgs := pgx.NamedArgs{
		"name":       user.Name,
		"username":   user.Username,
		"email":      user.Email,
		"password":   user.Password,
		"created_at": user.CreatedAt,
	}
	query := `
		INSERT INTO users (name, username, email, password, created_at)
		VALUES (@name, @username, @email, @password, @created_at)
	`
	_, err := database.DB.Exec(ctx, query, userArgs)

	if err != nil {
		return fmt.Errorf("insert user row is error: Unable to insert user row %v", err)
	}

	return nil
}

func (user User) SelectAll(ctx context.Context) ([]ResponseUser, error) {
	var users []ResponseUser
	query := `
		SELECT id, name, username, email, created_at, deleted_at
		FROM users
	`
	rows, err := database.DB.Query(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("select users is error: unable to get all users %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		user := ResponseUser{}
		err := rows.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.CreatedAt, &user.DeletedAt)
		if err != nil {
			return nil, fmt.Errorf("scan user is error: unable to get row %v", err)
		}
		users = append(users, user)
	}

	return users, nil
}

func (user *User) SelectByEmail(ctx context.Context, email string) error {
	userArgs := pgx.NamedArgs{
		"email": email,
	}
	query := `
		SELECT * 
		FROM users
		WHERE email = @email
	`
	err := database.DB.QueryRow(ctx, query, userArgs).Scan(
		&user.Id,
		&user.Name,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.DeletedAt,
	)

	if err != nil {
		return fmt.Errorf("scan select user by email is wrong: %v", err)
	}

	return nil
}
