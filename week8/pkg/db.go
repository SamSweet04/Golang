package pkg

import (
	"database/sql"	
	// _ "github.com/lib/pq"
)

func Db() *sql.DB{
	connStr := "user=postgres password=Haker15987 dbname=golang sslmode=disable"
    DB, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    } 
    // defer DB.Close()
	return DB
}