package db

import "database/sql"

type MyDb struct {
	Db *sql.DB
}

var instance *MyDb
func init() {

}

func Connect() {
		_db, err := sql.Open("mysql", "root:1234@/sqlc?parseTime=true")
		if err != nil {
			return
		}
		instance = &MyDb{Db: _db}
}

func GetConnection() MyDb {
	Connect()
	return *instance
}

func Close() {
	instance.Db.Close()
}