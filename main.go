package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/DAdDY0055/go_echo_firebase/models"

	"github.com/labstack/echo/v4"
	_ "github.com/go-sql-driver/mysql"
)


func main() {
	db, err := sql.Open("mysql", "root:pw@tcp(mysql)/db")
	fmt.Println("db err:", err)
	users, err := models.Users().All(context.Background(), db)
	fmt.Println("users err:", err)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		fmt.Println("Hello, Golang!")
		fmt.Println("users", users)
		return c.String(http.StatusOK, "Hello, Golang!2")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
