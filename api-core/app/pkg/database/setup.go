package database

import (
	"database/sql"
	"sync"

	"github.com/dnguyenngoc/doc-extractor/pkg/setting"
	_ "github.com/mattn/go-sqlite3"
)

var (
	dbPool     *sql.DB
	dbPoolOnce sync.Once
)

// GetDBConnection returns a connection from the database connection pool
func GetDBConnection() *sql.DB {
	dbPoolOnce.Do(func() {
		_ = CreateDBPool()
	})
	return dbPool
}

// CreateDBPool creates the database connection pool
func CreateDBPool() error {
	connectionString := setting.Config.GetString("XtractDbUri")
	db, err := sql.Open("sqlite3", connectionString)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	// Set the maximum number of open connections in the pool
	db.SetMaxOpenConns(10)

	// Set the maximum number of idle connections in the pool
	db.SetMaxIdleConns(5)

	// Assign the created connection pool to the dbPool variable
	dbPool = db

	return nil
}

// CloseDBPool closes the database connection pool
func CloseDBPool() error {
	if dbPool != nil {
		err := dbPool.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
