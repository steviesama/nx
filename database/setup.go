package database

import (
  "fmt"
  "database/sql"
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
    fmt.Println("MySQL Connection Closed!")
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
  db, err = sql.Open("mysql", connString)

  if err != nil {
    fmt.Printf("MySQL: %v\n", err)
    return nil
  } else {
    fmt.Println("MySQL connection successful!")
  }

  return db
}
