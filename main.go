package main

import (
	"fmt"
	d "urlpro/initializers/database"
	r "urlpro/initializers/routes"
)

func main() {

	//exampleQuery := `select * from urlpro.urls`
	//r, e := db.Exec(exampleQuery)
	//
	//fmt.Println("Result is >> ", r)

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
		fmt.Printf("originalUrl: %s\n", *expirydate)

	}

	//fmt.Println(con)

	// InitRoutes initiated gin and declares all the routes required
	// for the application.
	r.InitRoutes()

}
