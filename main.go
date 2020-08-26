package main

import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)



func main() {
	fmt.Println("ciao")

	db, err := sql.Open("mysql", "test_golang:test_go_126@tcp(127.0.0.1:3306)/GolangTest")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println("Succesfully connected to my db")

	insert, err := db.Query("INSERT INTO testUser(username, password) VALUES('Andrea', 'testPassword123')")

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	fmt.Println("Succesfully created a new User")

	

}

