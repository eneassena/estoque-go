package dbmysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type IDataBase struct{}

// type IDataBaseOperations interface {
// 	Query(query string, args ...any) (*sql.Rows, error)
// 	QueryRow(query string, args ...any) *sql.Row
// 	Prepare(query string, args ...any) (sql.Result, error)
// 	Exec(query string, args ...any) (sql.Result, error)
// }

func (*IDataBase) Open() *sql.DB {
	user := os.Getenv("DATABASE_USERNAME")
	pass := os.Getenv("DATABASE_PASSWORD")
	port := os.Getenv("DATABASE_PORT")
	host := os.Getenv("DATABASE_HOST")
	database := os.Getenv("DATABASE_NAME")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, database)
	log.Println("dataSource: ", dataSource)
	db, err := sql.Open("mysql", dataSource)

	if err != nil {
		log.Fatal(err.Error())
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err.Error())
	}
	return db
}
