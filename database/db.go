package database

import (
	"database/sql"
	"github.com/DavidSkeppstedt/blog/credentialsutil"
	_ "github.com/lib/pq"
	"log"
)

func Open() (db *sql.DB, err error) {
	dbConf, err := credentialsutil.CreateConfig("config/db.conf")
	if err != nil {
		log.Println("Error with database config", err.Error())
		panic(err)
	}
	db, err = sql.Open("postgres", dbConf)

	if err != nil {
		log.Println("Problem with the database driver", err.Error())
		panic(err)
	}
	return db, err
}
