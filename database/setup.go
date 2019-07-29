package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// registerConnectionPool initializes a connection pool pointer in the dbs
// map with the passed pool key.
// It returns a boolean indicating whether it was successful.
func registerConnectionPool(poolKey string) bool {
	_, ok := dbs[poolKey]
	// poolKey exists...return failure
	if ok {
		return false
	}

	dbs[poolKey] = nil

	return true
}

// connect establishes a connection to the specified database using the connInfo
// passed to it.
// It returns nil if a connection wasn't established, or a reference to the pool
// if it was.
func connect(connInfo ConnectionInfo) *sql.DB {
	var db *sql.DB
	// if the db has ben opened before, close it
	if db != nil {
		db.Close()
		fmt.Printf("%s Connection Closed!", connInfo.DbIdentity)
	}

	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&timeout=5s",
		connInfo.Username,
		connInfo.Password,
		connInfo.Url,
		connInfo.Port,
		connInfo.DbName,
	)

	var err error
	db, err = sql.Open(string(connInfo.DbIdentity), connString)

	if err != nil {
		fmt.Printf("%s: %v\n", connInfo.DbIdentity, err)
		return nil
	} else {
		fmt.Printf("%s connection successful!\n", connInfo.DbIdentity)
	}

	return db
}
