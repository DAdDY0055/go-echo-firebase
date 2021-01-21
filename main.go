package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/DAdDY0055/go_echo_firebase/models"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
	"github.com/labstack/echo/v4"
	_ "github.com/go-sql-driver/mysql"
)


func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql)/echo")
	fmt.Println("db err:", err)
	users, err := models.Users().All(context.Background(), db)
	fmt.Println("users err:", err)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		fmt.Println("Hello, Golang!")
		fmt.Println("users", users)
		return c.String(http.StatusOK, "Hello, Golang!2")
	})
	e.GET("/users", registerUser) // TODO: 一旦ブラウザで叩くためGET
	e.Logger.Fatal(e.Start(":8080"))
}

func registerUser(c echo.Context) error {
	ctx := context.Background()
	app, _ := initFirebaseApp()

	client, err := app.Auth(ctx)
	if err != nil {
					log.Fatalf("error getting Auth client: %v\n", err)
	}

	params := (&auth.UserToCreate{}).
		Email("user@example.com").
		EmailVerified(false).
		PhoneNumber("+15555550100").
		Password("secretPassword").
		DisplayName("John Doe").
		PhotoURL("http://www.example.com/12345678/photo.png").
		Disabled(false)
	u, err := client.CreateUser(ctx, params)
	if err != nil {
		log.Fatalf("error creating user: %v\n", err)
	}
	fmt.Println("u", u)
	return c.String(http.StatusOK, "ユーザーを作成しました")
}

func initFirebaseApp() (*firebase.App, error) {
	opt := option.WithCredentialsFile("cred/firebase_sec_key.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return app, err
}
