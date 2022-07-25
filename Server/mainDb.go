package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	// "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/michzuerch/CheckInBoard/models"
)




func main() {
  fmt.Println("Starting...")
	db, err := sql.Open("postgres", "dbname=CheckInBoard host=localhost user=michzuerch sslmode=disable")


	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Printf("Database: %v\n", db)

	ctx := context.Background()

	fmt.Printf("Context: %v",ctx)

	persons, err := models.Persons().All(ctx, db)


	if err != nil {
		fmt.Printf("Error persons: %v", err)
	}

	fmt.Printf("Persons %v\n",persons)

	for _, person := range persons {

		fmt.Printf("Person info: %v\n", person)
	}
	fmt.Println("End...")
}
