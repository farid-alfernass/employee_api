package repository

import (
	"errors"
	"log"
	"time"

	"github.com/faridtriwicaksono/employee_api/infrastructure/db/mysql"
	"github.com/faridtriwicaksono/employee_api/lib/converter"
	"github.com/faridtriwicaksono/employee_api/models"
	"github.com/faridtriwicaksono/employee_api/services/employee"
)

type (
	mysqlEmployeeRepository struct {
		dbAccess mysql.DBInterface
	}
)

// NewMysqlEmployeeRepository will create an object that represent the employee.Repository interface
func NewMysqlEmployeeRepository(dbAccess mysql.DBInterface) employee.DBRepository {
	return &mysqlEmployeeRepository{dbAccess: dbAccess}
}

// GetNewEmployee Query
func (m mysqlEmployeeRepository) GetNewEmployee(perPage int, offset int) ([]models.Employee, error) {
	var result []models.Employee

	db, err := m.dbAccess.GetDB("employee")
	if err != nil {
		log.Println("Error employee/repository function GetNewEmployee", err)
		return result, err
	}

	rows, err := db.Queryx("Call GetAllEmployees(?, ?)", perPage, offset)
	if err != nil {
		log.Println("Failed to do GetAllEmployees() query ", err)
		return result, err
	}

	defer rows.Close()
	for rows.Next() {
		employee := models.Employee{}
		if err = rows.StructScan(&employee); err != nil {
			log.Println("Error scanning GetAllEmployees()", err)
			continue
		}

		result = append(result, employee)
	}

	return result, nil
}

// GetBySearchKeywords Query
func (m mysqlEmployeeRepository) GetBySearchKeywords(keywords string, perPage int, offset int) ([]models.Employee, error) {
	var result []models.Employee

	db, err := m.dbAccess.GetDB("employee")
	if err != nil {
		log.Println("Error employee/repository function GetBySearchKeywords", err)
		return result, err
	}

	rows, err := db.Queryx("Call GetEmployeesBySearch(?, ?, ?)", keywords, perPage, offset)
	if err != nil {
		log.Println("Failed to do GetEmployeesBySearch() query ", err)
		return result, err
	}

	defer rows.Close()
	for rows.Next() {
		employee := models.Employee{}
		if err = rows.StructScan(&employee); err != nil {
			log.Println("Error scanning GetEmployeesBySearch()", err)
			continue
		}

		result = append(result, employee)
	}

	return result, nil
}

// GetByID Query
func (m mysqlEmployeeRepository) GetByID(id string, utype string) (models.Employee, error) {
	var result models.Employee

	if id == "" {
		return result, errors.New("Invalid Employee ID")
	}

	db, err := m.dbAccess.GetDB("employee")
	if err != nil {
		log.Println("Error employee/repository function GetByID", err)
		return result, err
	}

	err = db.QueryRowx("Call GetEmployeeDetail(?, ?)", id, utype).StructScan(&result)
	if err != nil {
		log.Println("Failed to do GetEmployeeDetail() query ", err)
		return result, err
	}

	return result, nil
}

// InsertEmployee Query
func (m mysqlEmployeeRepository) InsertEmployee(employee *models.Employee) (models.Employee, error) {
	var result models.Employee

	if employee == nil {
		log.Println("Error employee/repository function InsertEmployee employee struct is nil")
		return result, errors.New("employee struct is nil")
	}

	db, err := m.dbAccess.GetDB("employee")
	if err != nil {
		log.Println("Error employee/repository function InsertEmployee", err)
		return result, err
	}

	tx, err := db.Beginx()
	if err != nil {
		log.Println("Error Beginx employee/repository function InsertEmployee", err)
		return result, err
	}

	query := `
        INSERT INTO employees (
			email,
			fullname,
			created_at
		)
        VALUES
        (?, ?, ?)`

	sqlResult, err := tx.Exec(query, employee.Email, employee.Fullname, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		tx.Rollback()
		log.Println("Error exec employee/repository function InsertEmployee", err, employee.ID)
		return result, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Println("Error commit employee/repository function InsertEmployee", err, employee.ID)
		return result, err
	}

	id, err := sqlResult.LastInsertId()
	if err != nil {
		log.Println("Error get LastInsertId function InsertEmployee", err)
	}

	result.ID = converter.ConvertInt(id)

	return result, nil
}

// UpdateEmployee Query
func (m mysqlEmployeeRepository) UpdateEmployee(employee *models.Employee) (models.Employee, error) {
	var result models.Employee

	if employee == nil {
		log.Println("Error employee/repository function UpdateEmployee employee struct is nil")
		return result, errors.New("employee struct is nil")
	}

	db, err := m.dbAccess.GetDB("employee")
	if err != nil {
		log.Println("Error employee/repository function UpdateEmployee", err)
		return result, err
	}

	tx, err := db.Beginx()
	if err != nil {
		log.Println("Error Beginx employee/repository function UpdateEmployee", err)
		return result, err
	}

	query := `
        UPDATE employees SET
			email = ?,
			fullname = ?,
			updated_at = ?
		WHERE id = ?`

	_, err = tx.Exec(query, employee.Email, employee.Fullname, time.Now().Format("2006-01-02 15:04:05"), employee.ID)
	if err != nil {
		tx.Rollback()
		log.Println("Error exec employee/repository function UpdateEmployee", err, employee.ID)
		return result, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Println("Error commit employee/repository function UpdateEmployee", err, employee.ID)
		return result, err
	}

	return result, nil
}

// DeleteEmployee Query
func (m mysqlEmployeeRepository) DeleteEmployee(employee *models.Employee) (models.Employee, error) {
	var result models.Employee

	if employee == nil {
		log.Println("Error employee/repository function DeleteEmployee employee struct is nil")
		return result, errors.New("employee struct is nil")
	}

	db, err := m.dbAccess.GetDB("employee")
	if err != nil {
		log.Println("Error employee/repository function DeleteEmployee", err)
		return result, err
	}

	tx, err := db.Beginx()
	if err != nil {
		log.Println("Error Beginx employee/repository function DeleteEmployee", err)
		return result, err
	}

	query := `
        UPDATE employees SET deleted_at = ?
		WHERE id = ?`

	_, err = tx.Exec(query, time.Now().Format("2006-01-02 15:04:05"), employee.ID)
	if err != nil {
		tx.Rollback()
		log.Println("Error exec employee/repository function DeleteEmployee", err, employee.ID)
		return result, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Println("Error commit employee/repository function DeleteEmployee", err, employee.ID)
		return result, err
	}

	return result, nil
}

// GetEmployee Query
func (m mysqlEmployeeRepository) GetEmployee(employeeID string) (int, error) {
	var id int

	db, err := m.dbAccess.GetDB("employee")
	if err != nil {
		log.Println("Error employee/repository function GetEmployee", err)
		return id, err
	}

	err = db.QueryRowx("SELECT id from employees WHERE fullname = ?", employeeID).Scan(&id)
	if err != nil {
		log.Println("Failed to do GetEmployee() query by fullname ", err)
		if err.Error() == "sql: no rows in result set" {
			err = db.QueryRowx("SELECT id from employees WHERE id = ?", employeeID).Scan(&id)
			if err != nil {
				log.Println("Failed to do GetEmployee() query by id ", err)
				return id, err
			}
		}
	}
	return id, nil
}
