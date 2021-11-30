package main

import (
	"context"
	"fmt"
	"log"
	"time"

	//"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgxpool"
	"os"
)

var dbPool *pgxpool.Pool

func main() {

	// get the database connection URL.
	// usually, this is taken as an environment variable as in below commented out code
	// databaseUrl = os.Getenv("DATABASE_URL")

	// for the time being, let's hard code it as follows.
	// ensure to change values as needed.
	databaseUrl := "postgres://postgres:111@localhost:5432/Test"

	// this returns connection pool
	var err error
	dbPool, err = pgxpool.Connect(context.Background(), databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	Query()

	// to close DB pool
	defer dbPool.Close()

}

func Query() {

	// execute the select query and get result rows
	rows, err := dbPool.Query(context.Background(), "select * from \"Users\"")
	if err != nil {
		log.Fatal("error while executing query")
	}

	// iterate through the rows
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}

		// convert DB types to Go types
		firstName := values[0].(string)
		sdateOfBirth := ""
		if values[1] != nil {
			dateOfBirth := values[1].(time.Time)
			sdateOfBirth = (dateOfBirth.String())
		}
		id := values[2].(int32)

		log.Println("[id:", id, ", first_name:", firstName, ", date_of_birth:", sdateOfBirth, "]")
	}

}
