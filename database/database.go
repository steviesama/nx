// nx/database provides a way to manage database connections via key/value.
// The caller can setup a connection with a key...and provide other packages
// access to the database connection pool instance via dependency injection.
package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Identity represents the name of the rdbms that should be used when
// it is associated with an action. Such as "mysql".
type Identity string

const (
	// Used to represent the MySQL database driver.
	MysqlIdentity Identity = "mysql"
)

var (
	// a key/value map of connection pools i.e. key/pool
	dbs map[string]*sql.DB
)

func init() {
	// make the dbs map
	dbs = make(map[string]*sql.DB)
}

// New creates a new connection pool identified by poolKey and established by
// connInfo. Once built, if the connection is good, it sets the
// max idle & open connections based on the maxIdleConns & maxOpenConns
// parameters.
// It returns nil if the connection pool couldn't be built, or a reference
// to it if it could.
func New(poolKey string, connInfo ConnectionInfo) *sql.DB {
	isGoodKey := registerConnectionPool(poolKey)

	if !isGoodKey {
		fmt.Println("isGoodKey == false")
		return nil
	}

	dbs[poolKey] = connect(connInfo)

	if dbs[poolKey] != nil {
		dbs[poolKey].SetMaxIdleConns(connInfo.MaxIdleConns)
		dbs[poolKey].SetMaxOpenConns(connInfo.MaxOpenConns)
	}

	return dbs[poolKey]
}

// Delete closes and removes the connection pool associated with the pool key
// parameter.
// It returns a boolean indicating whether or not it the pool key existed. If it
// does...it will have deleted it.
func Delete(poolKey string) bool {
	_, ok := dbs[poolKey]
	if !ok {
		return false
	}

	dbs[poolKey].Close()
	delete(dbs, poolKey)

	return true
}

// DeleteAll iterates over all database pool pointers and attempts to close and
// delete them.
// It returns false if any close/delete fails...and true if all succeed.
func DeleteAll() bool {
	success := true
	for key := range dbs {
		if Delete(key) == false {
			success = false
		}
	}

	return success
}

// Pool takes a pool key to reference a database connection pool in the dbs
// connection pool map.
// It returns nil if no such key exists...or a reference to the pool if it does.
func Pool(poolKey string) *sql.DB {
	pool, ok := dbs[poolKey]

	if !ok {
		return nil
	}

	return pool
}
