package service

import (
	"tutorial.sqlc.dev/app/db"
	"tutorial.sqlc.dev/app/model"
)

var (
	dao *model.Queries
)

func Init() {
	dao = model.New(db.GetConnection().Db)
}
