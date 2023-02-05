package database

import (
	"fmt"
)

func SetupDatabase(){
	db := Connection()

	tableQueries := TablesSetup()

	for i := 0; i < len(tableQueries); i++ {
		_, err := db.Exec(tableQueries[i])

		if err != nil {
			panic(err)
		}

		fmt.Println("Created table with query:", tableQueries[i])
	}
}