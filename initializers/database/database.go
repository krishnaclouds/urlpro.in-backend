package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	host     = getEnvValue("DB_HOST")
	port     = 5432
	user     = getEnvValue("DB_USER")
	password = getEnvValue("DB_PASSWORD")
	dbname   = getEnvValue("DB_NAME")
)

func getEnvValue(key string) string {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	value, ok := viper.Get(key).(string)
	if ok {
		return value
	} else {
		panic("Unable to fetch Config")
	}
}

// ConnectDB is used to Connect to Database
func ConnectDB() (*sql.DB, error) {

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	fmt.Println(db)
	CheckError(err)

	// TODO: Figure out a way to close the db properly
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
