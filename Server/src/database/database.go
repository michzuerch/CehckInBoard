package database

import (
	"database/sql"
	"log"
	"os"

	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {

	db, err := sql.Open("postgres", "dbname=CheckInData user=michzuerch")
	if err != nil {
		log.Fatal("Failed to connect to the database.\n", err.Error())
		os.Exit(2)
	}

	log.Println("Connected to the database successfuly")

}
