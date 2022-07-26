package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	// "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/michzuerch/CheckInBoard/models"
)




func main() {
  fmt.Println("Starting...")
	db, err := sql.Open("postgres", "dbname=CheckInBoard host=localhost user=michzuerch sslmode=disable")


	if err != nil {
		panic(err)
	}

	ctx := context.Background() 

	persons, err := models.Persons().All(ctx, db)

	if err != nil {
		fmt.Printf("Error read persons from database: %v\n",err)
		panic(err)
	}

	// fmt.Printf("Persons %v\n",persons)

	currentTime := time.Now()

//Create new person
	var p models.Person
	p.Firstname = "Felix"
	p.Lastname = "Tester"
	p.Emailaddress = null.StringFrom("felix@unknowm.ee")
  p.Birthdate = null.TimeFrom(currentTime) 

  err2 := p.Insert(ctx,db, boil.Infer())

	if err2 != nil {
		fmt.Printf("Error insert to database: %v\n", err2)
		panic(err)
	}

	for _, person := range persons {

		fmt.Printf("Person firstname: %v\n", person.Firstname)
	}
	fmt.Println("End...")
}
