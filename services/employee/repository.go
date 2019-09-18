package employee

import (
	"github.com/faridtriwicaksono/employee_api/models"
)

// DBRepository represent the employee's repository contract
type DBRepository interface {
	GetNewEmployee(perPage int, offset int) ([]models.Employee, error)
	GetBySearchKeywords(keywords string, perPage int, offset int) ([]models.Employee, error)
	GetByID(id string, utype string) (models.Employee, error)
	InsertEmployee(user *models.Employee) (models.Employee, error)
	UpdateEmployee(user *models.Employee) (models.Employee, error)
	DeleteEmployee(user *models.Employee) (models.Employee, error)
	GetEmployee(userID string) (int, error)
}

// CacheRepository represent the employee's cache repository contract
type CacheRepository interface {
	GetNewEmployee(perPage int, offset int) ([]models.Employee, error)
	SetNewEmployee(perPage int, offset int, val []models.Employee) error

	GetBySearchKeywords(keywords string, perPage int, offset int) ([]models.Employee, error)
	SetBySearchKeywords(keywords string, perPage int, offset int, val []models.Employee) error

	GetByID(id string, utype string) (models.Employee, error)
	SetByID(id string, utype string, val models.Employee) error
}
