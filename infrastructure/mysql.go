package infrastructure

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// MySQLDatabase to connect to MySQL
func MySQLDatabase() *gorm.DB {
	db, err := gorm.Open("mysql", "root:farid123pir/employee?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("Error connecting to DB: %s", err)
	}

	db.DB().SetMaxIdleConns(0)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Minute * 10)

	return db
}
