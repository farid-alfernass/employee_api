package models

import (
	"time"
)

// User struct
type (
	Employee struct {
		ID               int        `json:"employee_id"`
		Email         	 string     `json:"email" db:"email"`
		Fullname         string     `json:"fullname" db:"fullname"`
		CreatedAt        *time.Time `json:"created_at" db:"created_at"`
		UpdatedAt        *time.Time `json:"updated_at" db:"updated_at"`
		DeletedAt        *time.Time `json:"deleted_at" db:"deleted_at"`
	}
)
