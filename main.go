package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/DAdDY0055/go_echo_firebase/models"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"github.com/labstack/echo/v4"
	_ "github.com/go-sql-driver/mysql"
)


func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql)/echo")
	fmt.Println("db err:", err)
	users, err := models.Users().All(context.Background(), db)
	fmt.Println("users err:", err)

	opt := option.WithCredentialsFile("cred/firebase_auth.json")
	config := &firebase.Config{ProjectID: "echo-api-2a307"}
	_, err = firebase.NewApp(context.Background(), config, opt) // TODO: appとして利用
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		fmt.Println("Hello, Golang!")
		fmt.Println("users", users)
		return c.String(http.StatusOK, "Hello, Golang!2")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
