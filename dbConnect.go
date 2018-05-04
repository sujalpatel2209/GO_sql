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
	selDb, err := db.Query("SELECT * FROM user")
	user := User{}
	res := []User{}
	if err != nil{
		panic(err.Error())
	}
	for selDb.Next() {
		var id, age int
		var firstname, lastname string
		err := selDb.Scan(&id, &firstname, &lastname, &age)
		if err != nil{
			panic(err.Error())
		}
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
	if err != nil{
		panic(err.Error())
	}
	return db
}

