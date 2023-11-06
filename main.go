package main

import (
	"fmt"
	"github.com/spf13/viper"

	d "urlpro/initializers/database"
	r "urlpro/initializers/routes"
)

func main() {

	// Load Config
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	value, ok := viper.Get("APP_NAME").(string)

	if ok {
		fmt.Println(value)
	}

	// DB Example
	conn, err := d.ConnectDB()

	if err != nil {
		panic(err)
	}

	fmt.Println("Conn is >> ", conn)

	rows, err := conn.Query("select * from urlpro.urls")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			originalUrl  string
			shortUrlCode string
			baseUrl      *string
			expirydate   *string
		)

		if err := rows.Scan(&originalUrl, &shortUrlCode, &baseUrl, &expirydate); err != nil {
			panic(err)
		}

		fmt.Printf("originalUrl: %s\n", originalUrl)
		fmt.Printf("shortUrlCode: %s\n", shortUrlCode)
		fmt.Printf("baseUrl: %s\n", *baseUrl)

		if expirydate != nil {
			fmt.Printf("originalUrl: %s\n", *expirydate)
		}

	}

	//fmt.Println(con)

	// InitRoutes initiated gin and declares all the routes required
	// for the application.
	r.InitRoutes()

}
