package databaseAccess

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func DatabaseConn() *sql.DB {
	dsn := "root:123456@tcp(192.168.152.130:3306)/test"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil
	}
	return db
}
