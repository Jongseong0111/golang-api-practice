package service

import (
	"tutorial.sqlc.dev/app/db"
	"tutorial.sqlc.dev/app/tutorial"
)

var (
	dao *tutorial.Queries
)

func Init() {
	dao = tutorial.New(db.GetConnection().Db)
}
