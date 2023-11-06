package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "urlpro"
	password = "<PASSWORD>"
	dbname   = "urlpro"
)

// ConnectDB is used to Connect to Database
func ConnectDB() (*sql.DB, error) {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	fmt.Println(db)
	CheckError(err)

	//defer func(db *sql.DB) {
	//	err := db.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//}(db)

	return db, nil

	//exampleQuery := `select * from urlpro.urls`
	//r, e := db.Exec(exampleQuery)
	//
	//fmt.Println("Result is >> ", r)

	//CheckError(e)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
