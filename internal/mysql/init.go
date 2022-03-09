package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var DB *sql.DB

func InitCon() {
	usernameAndPassword := fmt.Sprint(os.Getenv("MYSQL_DB_USER")) + ":" + fmt.Sprint(os.Getenv("MYSQL_DB_PASSWORD"))
	hostName := "tcp(" + fmt.Sprint(os.Getenv("MYSQL_DB_HOST")) + ":" + fmt.Sprint(os.Getenv("MYSQL_DB_PORT")) + ")"
	urlConnection := usernameAndPassword + "@" + hostName + "/" + fmt.Sprint(os.Getenv("MYSQL_DB_DATABASE")) + "?charset=utf8&parseTime=true&loc=UTC"

	fmt.Printf("⇨ Connect MYSQL to Server %s ... \n", hostName)

	db, err := sql.Open(os.Getenv("MYSQL_DB_DRIVER"), urlConnection)
	if err != nil {
		log.Fatalf("⇨ %s Data source %s:%s , Failed : %s \n", os.Getenv("MYSQL_DB_DRIVER"), os.Getenv("MYSQL_DB_HOST"), os.Getenv("MYSQL_DB_PORT"), err.Error())
	}

	fmt.Printf("⇨ %s Data source %s:%s , Successfully connected! \n", os.Getenv("MYSQL_DB_DRIVER"), os.Getenv("MYSQL_DB_HOST"), os.Getenv("MYSQL_DB_PORT"))

	DB = db
}

func Conn() (*sql.DB, error) {
	usernameAndPassword := fmt.Sprint(os.Getenv("MYSQL_DB_USER")) + ":" + fmt.Sprint(os.Getenv("MYSQL_DB_PASSWORD"))
	hostName := "tcp(" + fmt.Sprint(os.Getenv("MYSQL_DB_HOST")) + ":" + fmt.Sprint(os.Getenv("MYSQL_DB_PORT")) + ")"
	urlConnection := usernameAndPassword + "@" + hostName + "/" + fmt.Sprint(os.Getenv("MYSQL_DB_DATABASE")) + "?charset=utf8&parseTime=true&loc=UTC"

	fmt.Printf("⇨ Connect MYSQL to Server %s ... \n", hostName)

	db, err := sql.Open(os.Getenv("MYSQL_DB_DRIVER"), urlConnection)
	if err != nil {
		log.Fatalf("⇨ %s Data source %s:%s , Failed : %s \n", os.Getenv("MYSQL_DB_DRIVER"), os.Getenv("MYSQL_DB_HOST"), os.Getenv("MYSQL_DB_PORT"), err.Error())
	}

	fmt.Printf("⇨ %s Data source %s:%s , Successfully connected! \n", os.Getenv("MYSQL_DB_DRIVER"), os.Getenv("MYSQL_DB_HOST"), os.Getenv("MYSQL_DB_PORT"))

	return db, nil
}
