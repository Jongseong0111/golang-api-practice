package main

import (
	"fmt"
	"github.com/KrishanBhalla/reminder"
	"github.com/KrishanBhalla/reminder/notify"
	"github.com/KrishanBhalla/reminder/schedule"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/slack-go/slack"
	"reflect"
	"time"
	"tutorial.sqlc.dev/app/db"
	"tutorial.sqlc.dev/app/domain/product"
	productService "tutorial.sqlc.dev/app/domain/product/service"
	"tutorial.sqlc.dev/app/domain/user"
	userService "tutorial.sqlc.dev/app/domain/user/service"
)

func main() {
	InitDb()

	app := fiber.New()
	api := slack.New("xoxb-2748040306852-2745820870403-5aGFXcSxGN9fnObVceoUe7WD")
	db.Connect()
	defer db.Close()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello Go Server")
	})

	s := "[utils/SendMessageToSlack]"
	a :=time.Now().Format("3:04")
	fmt.Println(reflect.TypeOf(a))

	format := time.Kitchen
	sh, _ := schedule.NewSchedule(format, "0:00AM", "Local")
	r := schedule.IntervalRepeater{
		Interval:  time.Second,
		NumTimes: 60,
	}
	r.Repeat(sh)
	notify.NewPushbullet()
	rem := reminder.Reminder{
		Schedule: sh,
		Notifier: &notify.Desktop{},
	}

	rem.Remind("")
	if time.Now().Format("3:04") == "11:56" {

		channelID, timestamp, err := api.PostMessage(
			"C02NMENBH1N",
			slack.MsgOptionText("alert! you must fix it!", false),
		)
		fmt.Println(channelID, timestamp)

		if err != nil {
			fmt.Printf("%s %v\n", s, err)
		}
		fmt.Printf("slack message post successfully %s at %s", channelID, timestamp)
	}
	user.Init(app)
	product.Init(app)
	app.Listen(":3000")
}

func InitDb() {
	userService.Init()
	productService.Init()
}