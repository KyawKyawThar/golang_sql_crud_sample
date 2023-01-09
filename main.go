package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {

	//Connect to db
	db, err := sql.Open("pgx", "host=localhost port=5432 dbname=testgo user=postgres password=secret")

	if err != nil {
		log.Fatal(fmt.Sprintf("Can't connect to db: %v\n", err))
	}

	defer db.Close()

	log.Println("Connect to DB")

	err = db.Ping()

	if err != nil {
		log.Println("Can't ping to the db")
	}

	log.Println("Successful ping to db")

	//get rows from table
	err = getRowFromTable(db)

	if err != nil {
		log.Println("Can't get Row From Table")
	}

	//Insert a Row
	query := `insert into users(first_name,last_name) values ($1,$2)`

	_, err = db.Exec(query, "test1 first", "test1 last")

	if err != nil {
		log.Fatal("err")
	}

	fmt.Println("Insert into database")
	err = getRowFromTable(db)

	if err != nil {
		log.Println("Can't get Row From Table", err)
	}

	//Update a Row
	upd := `update users set first_name=$1,last_name=$2 where id=$3`
	_, err = db.Exec(upd, "thuhtet", "naing", 3)
	if err != nil {
		log.Println("Can't get Row From Table", err)
	}
	fmt.Println("Update into database")
	err = getRowFromTable(db)

	//Get one Row By ID
	query = `select id,first_name,last_name from users where id=$1`

	var firstName, lastName string
	var id int

	row := db.QueryRow(query, 3)

	err = row.Scan(&id, &firstName, &lastName)

	if err != nil {
		log.Fatal(err)

	}

	fmt.Println("Query Row returns", firstName, lastName, id)

	//delete Row

	query = `delete from users where id=$1`

	_, err = db.Exec(query, 3)

	if err != nil {
		log.Fatal(err)

	}

	fmt.Println("Delete into database")
	err = getRowFromTable(db)
}

func getRowFromTable(db *sql.DB) error {

	rows, err := db.Query("select id, first_name ,last_name from users")

	if err != nil {
		log.Println(err)
		return err
	}

	defer rows.Close()

	var firstName, lastName string
	var id int

	for rows.Next() {

		err := rows.Scan(&id, &firstName, &lastName)

		if err != nil {
			log.Println(err)
			return err
		}

		fmt.Println("Record is row is: ", id, firstName, lastName)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error scanning rows", err)
	}

	fmt.Println("-----------------------------")
	return nil
}
