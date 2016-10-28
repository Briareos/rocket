package sql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewConnection(username, password, host, port, dbName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", username, password, host, port, dbName))
	if err != nil {
		return nil, fmt.Errorf("open mysql connection: %v", err)
	}

	return db, nil
}
