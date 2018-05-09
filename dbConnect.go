package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

type User struct {
	id int `json:"id,omitempty"`
	firstname string `json:"firstname,omitempty"`
	lastname string `json:"lastname,omitempty"`
	age int `json:"age,omitempty"`
} 

func main() {
	db := dbConn()

	intDb, err := db.Prepare("INSERT user SET firstname=?, lastname=?, age=?")
	checkErr(err)

	response, err := intDb.Exec("Test","User",25)
	checkErr(err)

	id, err := response.LastInsertId();

	fmt.Println(id);

	selDb, err := db.Query("SELECT * FROM user")
	user := User{}
	res := []User{}
	checkErr(err)

	for selDb.Next() {
		var id, age int
		var firstname, lastname string
		err := selDb.Scan(&id, &firstname, &lastname, &age)
		checkErr(err)

		user.id = id
		user.firstname = firstname
		user.lastname = lastname
		user.age = age
		res = append(res, user)
	}
	fmt.Println("Data : ", res)
	defer db.Close()
}

func dbConn() (db *sql.DB){
	db, err := sql.Open("mysql","root:@/gocrud")
	checkErr(err)
	return db
}

func checkErr(err error){
	if err != nil{
		panic(err)
	}
}
