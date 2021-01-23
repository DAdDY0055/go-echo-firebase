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
	e.GET("/users/:Id", getUserByID)
	e.GET("/users", listUser)
	e.POST("/users", registerUser)
	e.PUT("/users/:Id", updateUser)
	e.DELETE("/users/:Id", deleteUserByID)

	e.Logger.Fatal(e.Start(":8080"))
}

// Firebase関連 TODO: 別モジュールに移動させる

// registerUser ユーザー登録
func registerUser(c echo.Context) error {
	ctx := context.Background()
	client, err := initFirebaseClient(ctx)
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

// getUserByID IDによるユーザー取得
func getUserByID(c echo.Context) error {
	ctx := context.Background()
	client, err := initFirebaseClient(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	uid := c.Param("Id")
	u, err := client.GetUser(ctx, uid)
	if err != nil {
		log.Fatalf("error getting user %s: %v\n", uid, err)
	}
	return c.String(http.StatusOK, u.Email)
}

// listUser 一覧によるユーザー取得
func listUser(c echo.Context) error {
	ctx := context.Background()
	client, err := initFirebaseClient(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	uid := c.Param("id")
	users := client.Users(ctx, "")
	if err != nil {
		log.Fatalf("error getting user %s: %v\n", uid, err)
	}
	fmt.Println("u", users)

	return c.String(http.StatusOK, "ユーザー一覧取得") // TODO: 表示
}

// updateUser ユーザー更新
func updateUser(c echo.Context) error {
	ctx := context.Background()
	client, err := initFirebaseClient(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	params := (&auth.UserToUpdate{}).
		Email("update_user@example.com").
		EmailVerified(true).
		PhoneNumber("+15555550100").
		Password("newPassword").
		DisplayName("John Doe").
		PhotoURL("http://www.example.com/12345678/photo.png").
		Disabled(false)

	uid := c.Param("id")
	u, err := client.UpdateUser(ctx, uid, params)
	if err != nil {
		log.Fatalf("error updating user: %v\n", err)
	}
	fmt.Println("u", u)
	return c.String(http.StatusOK, "ユーザーを更新しました")
}

// deleteUserByID IDによるユーザー削除
func deleteUserByID(c echo.Context) error {
	ctx := context.Background()
	client, err := initFirebaseClient(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	uid := c.Param("Id")
	err = client.DeleteUser(ctx, uid)
	if err != nil {
		log.Fatalf("error deleting user %s: %v\n", uid, err)
	}
	return c.String(http.StatusOK, "ユーザーを削除しました")
}

// initFirebaseClient FirebaseClient初期化
func initFirebaseClient(ctx context.Context) (*auth.Client, error) {
	opt := option.WithCredentialsFile("cred/firebase_sec_key.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return app.Auth(ctx)
}
