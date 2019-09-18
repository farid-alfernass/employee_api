package employee

import (
	"github.com/gin-gonic/gin"
	"github.com/faridtriwicaksono/employee_api/models"
)

// Usecase represent the employee's repository contract
type Usecase interface {
	GetNewEmployee(c *gin.Context) ([]models.Employee, error)
	GetBySearchKeywords(c *gin.Context) ([]models.Employee, error)
	GetByID(c *gin.Context) (models.Employee, error)
	InsertEmployee(c *gin.Context) (models.Employee, error)
	UpdateEmployee(c *gin.Context) (models.Employee, error)
	DeleteEmployee(c *gin.Context) (models.Employee, error)
}
