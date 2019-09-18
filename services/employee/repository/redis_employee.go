package repository

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/faridtriwicaksono/employee_api/infrastructure/cache/redis"
	"github.com/faridtriwicaksono/employee_api/models"
	"github.com/faridtriwicaksono/employee_api/services/employee"
)

type (
	redisEmployeeRepository struct {
		redisAccess redis.Cache
	}
)

const (
	redisNewEmployee              = "employees:index:%d:%d"
	redisEmployeeByID             = "employees:detail:%s:%s"
	redisEmployeeBySearchKeywords = "employees:search:keywords:%s:%d:%d"
)

// NewRedisEmployeeRepository will create an object that represent the employee.CacheRepository interface
func NewRedisEmployeeRepository(redisAccess redis.Cache) employee.CacheRepository {
	return &redisEmployeeRepository{redisAccess: redisAccess}
}

// GetNewEmployee Cache
func (m redisEmployeeRepository) GetNewEmployee(perPage int, offset int) ([]models.Employee, error) {
	var result []models.Employee

	res, err := redis.GetRedisValue(m.redisAccess, fmt.Sprintf(redisNewEmployee, perPage, offset))
	if err != nil {
		return result, err
	}

	if res == nil {
		return result, err
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// SetBySearchKeywords Cache
func (m redisEmployeeRepository) SetNewEmployee(perPage int, offset int, val []models.Employee) error {
	resJSON, err := json.Marshal(val)
	if err != nil {
		return errors.Wrap(err, "failed to marshal JSON")
	}

	redis.SetRedisValue(m.redisAccess, fmt.Sprintf(redisNewEmployee, perPage, offset), resJSON, 300)
	return nil
}

// GetBySearchKeywords Cache
func (m redisEmployeeRepository) GetBySearchKeywords(keywords string, perPage int, offset int) ([]models.Employee, error) {
	var result []models.Employee

	res, err := redis.GetRedisValue(m.redisAccess, fmt.Sprintf(redisEmployeeBySearchKeywords, keywords, perPage, offset))
	if err != nil {
		return result, err
	}

	if res == nil {
		return result, err
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// SetBySearchKeywords Cache
func (m redisEmployeeRepository) SetBySearchKeywords(keywords string, perPage int, offset int, val []models.Employee) error {
	resJSON, err := json.Marshal(val)
	if err != nil {
		return errors.Wrap(err, "failed to marshal JSON")
	}

	redis.SetRedisValue(m.redisAccess, fmt.Sprintf(redisEmployeeBySearchKeywords, keywords, perPage, offset), resJSON, 300)
	return nil
}

// GetByID Cache
func (m redisEmployeeRepository) GetByID(id string, utype string) (models.Employee, error) {
	var result models.Employee

	res, err := redis.GetRedisValue(m.redisAccess, fmt.Sprintf(redisEmployeeByID, id, utype))
	if err != nil {
		return result, err
	}

	if res == nil {
		return result, err
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// SetByID Cache
func (m redisEmployeeRepository) SetByID(id string, utype string, val models.Employee) error {
	resJSON, err := json.Marshal(val)
	if err != nil {
		return errors.Wrap(err, "failed to marshal JSON")
	}

	redis.SetRedisValue(m.redisAccess, fmt.Sprintf(redisEmployeeByID, id, utype), resJSON, 300)
	return nil
}
