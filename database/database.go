
package database

import (
	"database/sql"
	//"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


func ConnectDatabase() (*sql.DB, error) {
	fmt.Println("Connecting to dn")
	db, err := sql.Open("mysql", "root:example@/post-api")
	if err != nil {
		fmt.Println("error from connect", err)
		return nil, err
	}

	return db, nil
}