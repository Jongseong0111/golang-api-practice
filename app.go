package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"reflect"
	"time"
	"tutorial.sqlc.dev/app/tutorial"

	_ "github.com/go-sql-driver/mysql"
)

func run() error {

	ctx := context.Background()

	db, err := sql.Open("mysql", "root:1234@/sqlc")
	if err != nil {
		return err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	fmt.Println("connect success", db)
	defer db.Close()

	queries := tutorial.New(db)

	// list all authors
	authors, err := queries.ListUser(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// create an author
	result, err := queries.CreateUser(ctx, tutorial.CreateUserParams{
		UserName:     "HKJA7",
		UserEmail:    "qkqndudnxld7@gmail.com",
		UserAccount:  "Leon7",
		UserPassword: "dudnxld",
	})
	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	log.Println(userID)

	// get the author we just inserted
	fetchedUser, err := queries.GetUser(ctx, int32(userID))
	if err != nil {
		return err
	}

	// prints true
	v1 := reflect.ValueOf(userID)
	v2 := reflect.ValueOf(fetchedUser.UserID)
	fmt.Println(v1.Type())
	fmt.Println(v2.Type())
	log.Println(reflect.DeepEqual(int32(userID), fetchedUser.UserID))
	return nil
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//if err := run(); err != nil {
	//	log.Fatal(err)
	//}
	type User struct {
		UserName     string `json:"userName"`
		UserAccount  string `json:"userAccount"`
		UserEmail    string `json:"userEmail"`
		UserPassword string `json:"userPassword"`
	}

	db, err := sql.Open("mysql", "root:1234@/sqlc")
	if err != nil {
		return
	}

	defer db.Close()

	queries := tutorial.New(db)
	app.Post("/user", func(ctx *fiber.Ctx) error {
		user := User{}
		if err := ctx.BodyParser(&user); err != nil {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}

		params := tutorial.CreateUserParams{
			UserName:  user.UserName,
			UserEmail: user.UserEmail,

			UserAccount:  user.UserAccount,
			UserPassword: user.UserPassword,
		}

		result, _ := queries.CreateUser(context.Background(), params)

		userID, err := result.LastInsertId()
		if err != nil {
			return err
		}

		fetchedUser, err := queries.GetUser(context.Background(), int32(userID))
		if err != nil {
			return err
		}

		return ctx.JSON(fetchedUser)
	})
	app.Listen(":3000")
}
