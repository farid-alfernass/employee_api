package usecase

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/faridtriwicaksono/employee_api/lib/converter"
	"github.com/faridtriwicaksono/employee_api/models"
	"github.com/faridtriwicaksono/employee_api/services/employee"
)

type userUsecase struct {
	db    employee.DBRepository
	cache employee.CacheRepository
}

// NewEmployeeUsecase will create new an userUsecase object representation of article.Usecase interface
func NewEmployeeUsecase(db employee.DBRepository, cache employee.CacheRepository) employee.Usecase {
	return &userUsecase{
		db:    db,
		cache: cache,
	}
}

// GetNewEmployee of employee function
func (uc userUsecase) GetNewEmployee(c *gin.Context) ([]models.Employee, error) {
	var (
		result                            []models.Employee
		page, perPage, maxPerPage, offset int = employee.DefaultPage, employee.DefaultPerPage, employee.DefaultMaxPerPage, employee.DefaultOffset
	)

	if c.Query("page") != "" {
		page = converter.ConvertInt(c.Query("page"))
	}

	if c.Query("per_page") != "" {
		perPage = converter.ConvertInt(c.Query("per_page"))

		if perPage > maxPerPage {
			perPage = maxPerPage
		}

		offset = perPage * (page - 1)
	}

	result, err := uc.cache.GetNewEmployee(perPage, offset)
	if err == redis.Nil {
		result, err = uc.db.GetNewEmployee(perPage, offset)
		if err != nil {
			return result, err
		}

		err = uc.cache.SetNewEmployee(perPage, offset, result)
		if err != nil {
			return result, err
		}
	} else if err != nil {
		return result, err
	}

	return result, nil
}

// GetBySearchKeywords of employee function
func (uc userUsecase) GetBySearchKeywords(c *gin.Context) ([]models.Employee, error) {
	var (
		result                            []models.Employee
		page, perPage, maxPerPage, offset int = employee.DefaultPage, employee.DefaultPerPage, employee.DefaultMaxPerPage, employee.DefaultOffset
		keywords                              = employee.DefaultSearchKeywords
	)

	if c.PostForm("page") != "" {
		page = converter.ConvertInt(c.PostForm("page"))
	}

	if c.PostForm("per_page") != "" {
		perPage = converter.ConvertInt(c.PostForm("per_page"))

		if perPage > maxPerPage {
			perPage = maxPerPage
		}

		offset = perPage * (page - 1)
	}

	if c.PostForm("keywords") != "" {
		keywords = c.PostForm("keywords")
	}

	result, err := uc.cache.GetBySearchKeywords(keywords, perPage, offset)
	if err == redis.Nil {
		result, err = uc.db.GetBySearchKeywords(keywords, perPage, offset)
		if err != nil {
			return result, err
		}

		err = uc.cache.SetBySearchKeywords(keywords, perPage, offset, result)
		if err != nil {
			return result, err
		}
	} else if err != nil {
		return result, err
	}

	return result, nil
}

// GetByID Get Employee By ID
func (uc userUsecase) GetByID(c *gin.Context) (models.Employee, error) {
	var (
		result models.Employee
		utype  = employee.DefaultEmployeeType
	)

	_, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utype = "string"
	}

	id := c.Param("id")

	result, err = uc.cache.GetByID(id, utype)
	if err == redis.Nil {
		result, err = uc.db.GetByID(id, utype)
		if err != nil {
			return result, err
		}

		err = uc.cache.SetByID(id, utype, result)
		if err != nil {
			return result, err
		}
	} else if err != nil {
		return result, err
	}

	return result, nil
}

// InsertEmployee of employee function
func (uc userUsecase) InsertEmployee(c *gin.Context) (models.Employee, error) {
	var (
		user models.Employee
	)

	err := c.BindJSON(&user)
	if err != nil {
		return user, err
	}

	result, err := uc.db.InsertEmployee(&user)
	if err != nil {
		return user, err
	}

	return result, nil
}

// UpdateEmployee of employee function
func (uc userUsecase) UpdateEmployee(c *gin.Context) (models.Employee, error) {
	var (
		res   models.Employee
		utype = employee.DefaultEmployeeType
	)

	id := c.Param("id")

	res, err := uc.db.GetByID(id, utype)
	if err != nil {
		return res, err
	}

	err = c.BindJSON(&res)
	if err != nil {
		return res, err
	}
	_, err = uc.db.UpdateEmployee(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

// DeleteEmployee of employee function
func (uc userUsecase) DeleteEmployee(c *gin.Context) (models.Employee, error) {
	var (
		res   models.Employee
		utype = employee.DefaultEmployeeType
	)

	id := c.Param("id")

	res, err := uc.db.GetByID(id, utype)
	if err != nil {
		return res, err
	}

	result, err := uc.db.DeleteEmployee(&res)
	if err != nil {
		return res, err
	}

	return result, nil
}
