package appcontext

import (
	"fmt"
	"log"

	"github.com/addressbook/services/server/config"
	"github.com/addressbook/services/server/helper"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type appContext struct {
	db     *sqlx.DB
	helper helper.QueryExecuter
}

var context *appContext

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func initDB() (*sqlx.DB, error) {

	fmt.Println(config.GetConnectionString())
	db, err := sqlx.Open("postgres", config.GetConnectionString())
	check(err)
	err = db.Ping()
	check(err)
	return db, nil
}

func InitContext() {
	db, _ := initDB()
	log.Print("Database Connection Established")

	context = &appContext{
		db: db,
		helper: &helper.QueryHelper{
			DB: db,
		},
	}
}

func GetDB() *sqlx.DB {
	return context.db
}
func GetHelper() helper.QueryExecuter {
	return context.helper
}
