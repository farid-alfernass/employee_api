package presenter

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/faridtriwicaksono/employee_api/services/employee"
)

type (
	// EmployeeHandler  represent the httphandler for employee
	EmployeeHandler struct {
		EmployeeUseCase employee.Usecase
	}
)

// NewEmployeeHTTPHandler represent handler
func NewEmployeeHTTPHandler(e *gin.Engine, us employee.Usecase) {
	handler := &EmployeeHandler{
		EmployeeUseCase: us,
	}

	r := e.Group("/v1")
	{
		router := r.Group("/employee")
		{
			router.GET("/", handler.GetNewEmployee)
			router.POST("/search", handler.GetBySearchKeywords)
			router.GET("/detail/:id", handler.GetByID)
			router.POST("/store", handler.InsertEmployee)
			router.PATCH("/update/:id", handler.UpdateEmployee)
			router.DELETE("/delete/:id", handler.DeleteEmployee)
		}
	}
}

// GetNewEmployee usecase
func (a EmployeeHandler) GetNewEmployee(c *gin.Context) {
	start := time.Now().Local()
	result, err := a.EmployeeUseCase.GetNewEmployee(c)
	if err != nil {
		c.JSON(500, gin.H{"status": 500, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": 200, "process_time": time.Since(start).String(), "data": result})
	return
}

// GetBySearchKeywords usecase
func (a EmployeeHandler) GetBySearchKeywords(c *gin.Context) {
	start := time.Now().Local()
	result, err := a.EmployeeUseCase.GetBySearchKeywords(c)
	if err != nil {
		c.JSON(500, gin.H{"status": 500, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": 200, "process_time": time.Since(start).String(), "data": result})
	return
}

// GetByID usecase
func (a EmployeeHandler) GetByID(c *gin.Context) {
	start := time.Now().Local()
	result, err := a.EmployeeUseCase.GetByID(c)
	if err != nil {
		c.JSON(500, gin.H{"status": 500, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": 200, "process_time": time.Since(start).String(), "data": result})
	return
}

// InsertEmployee usecase
func (a EmployeeHandler) InsertEmployee(c *gin.Context) {
	start := time.Now().Local()
	result, err := a.EmployeeUseCase.InsertEmployee(c)
	if err != nil {
		c.JSON(500, gin.H{"status": 500, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": 200, "process_time": time.Since(start).String(), "data": result})
	return
}

// UpdateEmployee usecase
func (a EmployeeHandler) UpdateEmployee(c *gin.Context) {
	start := time.Now().Local()
	result, err := a.EmployeeUseCase.UpdateEmployee(c)
	if err != nil {
		c.JSON(500, gin.H{"status": 500, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": 200, "process_time": time.Since(start).String(), "data": result})
	return
}

// DeleteEmployee usecase
func (a EmployeeHandler) DeleteEmployee(c *gin.Context) {
	start := time.Now().Local()
	result, err := a.EmployeeUseCase.DeleteEmployee(c)
	if err != nil {
		c.JSON(500, gin.H{"status": 500, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": 200, "process_time": time.Since(start).String(), "data": result})
	return
}
