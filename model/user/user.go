package user

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"govibes.app/database"
)

type User struct {
	Id        uuid.UUID   `json:"id"`
	Name      string      `json:"name"`
	Username  string      `json:"username"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
	CreatedAt time.Time   `json:"created_at"`
	DeletedAt pgtype.Date `json:"deleted_at"`
}

func (user *User) InsertUser(ctx context.Context, reqBody RequestRegister) error {
	userId, err := uuid.NewV7()

	if err != nil {
		return fmt.Errorf("error")
	}

	user.Id = userId
	user.Name = reqBody.Name
	user.Username = reqBody.Username
	user.Email = reqBody.Email
	user.Password = reqBody.Password
	user.CreatedAt = time.Now()

	userArgs := pgx.NamedArgs{
		"id":         user.Id,
		"name":       user.Name,
		"username":   user.Username,
		"email":      user.Email,
		"password":   user.Password,
		"created_at": user.CreatedAt,
	}

	query := `
		INSERT INTO users (id, name, username, email, password, created_at)
		VALUES (@id, @name, @username, @email, @password, @created_at)
	`
	_, err = database.DB.Exec(ctx, query, userArgs)

	if err != nil {
		return fmt.Errorf("insert user row is error: Unable to insert user row %v", err)
	}

	return nil
}

func (user User) SelectAll(ctx context.Context) ([]ResponseRegister, error) {
	var users []ResponseRegister
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
		user := ResponseRegister{}
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

func (user *User) UpdateRowNameById(ctx context.Context, reqBody RequestEditProfile) error {
	userArgs := pgx.NamedArgs{
		"id":   reqBody.Id,
		"name": reqBody.Name,
	}

	query := `
		UPDATE users
		SET name = @name
		WHERE id = @id
	`

	_, err := database.DB.Exec(ctx, query, userArgs)

	if err != nil {
		return fmt.Errorf("scan update is error: %v", err)
	}

	user.Id = reqBody.Id
	user.Name = reqBody.Name

	return nil
}
