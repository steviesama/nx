package database

import (
  "fmt"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

// TODO: this is unfinished...didn't need it...may need removal
func ScanType(v interface{}, rows *sql.Rows, colType *sql.ColumnType) error {
  fmt.Printf("DatabaseTypeName: %v\n", colType.DatabaseTypeName())
  return nil
}