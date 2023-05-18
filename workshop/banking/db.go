package banking

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
}

func (my *MySQL) CreateConnection() *sql.DB {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	return db
}
