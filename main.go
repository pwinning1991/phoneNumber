package main

import (
	"database/sql"
	"fmt"
	"os"
	"regexp"

	_ "github.com/lib/pq"
)

var (
	host     = "localhost"
	port     = 5432
	user     = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbname   = os.Getenv("POSTGRES_DB")
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = createDB(db, dbname)
	if err != nil {
		panic(err)
	}
	db.Close()
}

func resetDB(db *sql.DB, name string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		return err
	}
	return createDB(db, name)
}

func createDB(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE" + name)
	if err != nil {
		return err
	}
	return nil

}

func normalize(phone string) string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(phone, "")

}

//func normalize(phone string) string {
//	var buf bytes.Buffer
//
//	for _, ch := range phone {
//		if ch >= '0' && ch <= '9' {
//			buf.WriteRune(ch)
//		}
//	}
//	return buf.String()
//}
