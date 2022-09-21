package mariadb

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func New(url string) (*sql.DB, error) {
	conn, err := sql.Open("mysql", url)

	return conn, err
}
