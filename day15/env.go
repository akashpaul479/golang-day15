package day15

import (
	"fmt"
	"os"
)

func Env() {
	dbuser := os.Getenv("DB_USER")
	dbpass := os.Getenv("DB_PASS")

	if dbuser == "" || dbpass == "" {
		fmt.Println("Environment variables DB_USER or DB_PASS not set")
	} else {
		fmt.Printf("DB_USER=%s, DB_PASS=%s\n", dbuser, dbpass)
	}
}
