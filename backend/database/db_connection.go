package database

import (
	"database/sql"
	env "echo_chamber/env_variables"
	"fmt"

	_ "github.com/lib/pq"
)

func Connection() *sql.DB {
	config := env.ProjectEnvs()

	psqlAuth := fmt.Sprintf("host=%s port=%s user=%s " + 
	"password=%s dbname=%s sslmode=disable", config.Host, config.Port,
	 config.DBuser, config.Password, config.DBname)
	
	 db, err := sql.Open("postgres", psqlAuth)

	 if err != nil {
		fmt.Println(err)
		return nil
	 }

	 err = db.Ping()
	 
	 if err != nil {
		fmt.Println(err)
		return nil
	 }
	 
	 fmt.Println("Successfully connected to database!")

	 return db
}