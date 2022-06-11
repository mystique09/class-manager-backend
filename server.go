package main

import (
	"database/sql"
	"log"
	"server/app"
	"server/database/sqlc"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := sql.Open("postgres", "user=mystique09 password=mystique09 dbname=class-manager sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	qr := database.New(db)

	rt := app.Router{
		DB: qr,
	}

	config := app.Config{
		Port: ":8080",
	}

	server := echo.New()
	server.GET("/api/v1", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	server.GET("/api/v1/classes", rt.GetUsers)
	server.Logger.Fatal(server.Start(config.Port))
}
