package db

import (
	"sync"

	"github.com/faridtriwicaksono/employee_api/infrastructure/db/mysql"
)

var (
	accessOnce sync.Once
	access     mysql.DBInterface
)

// GetDBAccess provide new DB conn
func GetDBAccess() mysql.DBInterface {
	if access != nil {
		return access
	}

	accessOnce.Do(func() {
		dbClient := mysql.NewDatabase("root:farid123@tcp(127.0.0.1:3306)/employee?parseTime=true&loc=Local")
		dbClient.Connect("employee")
		access = dbClient
	})

	return access
}
