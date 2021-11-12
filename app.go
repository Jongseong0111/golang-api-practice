package main

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	"tutorial.sqlc.dev/app/tutorial"

	_ "github.com/go-sql-driver/mysql"
)

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

		duplicateAccount, err := queries.CheckDuplicateAccount(context.Background(), user.UserAccount)
		if err != nil {
			return err
		}

		if len(duplicateAccount) > 0 {
			err = errors.New("duplicate Account")
			return err
		}

		duplicateEmail, err := queries.CheckDuplicateEmail(context.Background(), user.UserEmail)
		if err != nil {
			return err
		}

		if len(duplicateEmail) > 0 {
			err = errors.New("duplicate Email")
			return err
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
