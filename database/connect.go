package database

import (
	"context"
	"fmt"

	"govibes.app/config"

	"github.com/jackc/pgx/v5"
)

func Connect()  {
	var err error
	databaseUrl := config.Config(config.DATABASE_URL_KEY)
	DB, err = pgx.Connect(context.Background(), databaseUrl)

	if err != nil {
		panic("Failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
}