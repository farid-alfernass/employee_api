package mysql

import (
	"log"
	"runtime"
	"time"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// DBConn variable to declare Database Connection
var DBConn *DatabaseConnection

type (
	// DBInterface to provide general Func
	DBInterface interface {
		Connect(string) *DatabaseConnection
		GetDB(string) (*sqlx.DB, error)
	}

	// Database to provide Database Config
	Database struct {
		Name DatabaseConfig
	}

	// DatabaseConfig currently have master only
	DatabaseConfig struct {
		Master string
	}

	// DatabaseConnection provide struct sqlx connection
	DatabaseConnection struct {
		Connection *sqlx.DB
	}
)

// NewDatabase create Database Struct from Config
func NewDatabase(config interface{}) *Database {
	cfg := config.(string)

	dc := DatabaseConfig{
		Master: cfg,
	}

	return &Database{
		Name: dc,
	}
}

// Connect provide sqlx connection
func (db *Database) Connect(dbName string) *DatabaseConnection {

	databaseConn := DatabaseConnection{}

	master := db.Name.Master
	if master != "" {
		db, err := sqlx.Connect("mysql", master)
		if err != nil {
			log.Println("Can not connect MySQL ", err)
		}

		db.SetMaxOpenConns(100)
		db.SetMaxIdleConns(10)
		db.SetConnMaxLifetime(time.Minute * 10)

		databaseConn.Connection = db
	}

	DBConn = &DatabaseConnection{Connection: databaseConn.Connection}

	return DBConn
}

// GetDB provide status from DB
func (db *Database) GetDB(dbName string) (*sqlx.DB, error) {
	var newDB *sqlx.DB

	newDB = DBConn.Connection

	if newDB.Stats().OpenConnections > 40 {
		fpcs := make([]uintptr, 1)
		n := runtime.Callers(2, fpcs)
		if n != 0 {
			fun := runtime.FuncForPC(fpcs[0] - 1)
			if fun != nil {
				log.Println("Db Conn more than 40, Caller from Func :", fun.Name())
			}
		}

		log.Println("DB Conn more than 40, currently :", newDB.Stats())
	}

	return newDB, nil
}
