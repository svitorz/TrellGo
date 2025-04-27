package models

import (
	"database/sql/driver"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Status string

const (
	StatusPendente    Status = "pendente"
	StatusEmAndamento Status = "em_andamento"
	StatusConcluida   Status = "concluida"
)

func (s Status) IsValid() bool {
	switch s {
	case StatusPendente, StatusEmAndamento, StatusConcluida:
		return true
	}
	return false
}

func (s *Status) Scan(value interface{}) error {
	strValue, ok := value.(string)
	if !ok {
		return errors.New("tipo incompatível para Status")
	}

	*s = Status(strValue)
	if !s.IsValid() {
		return fmt.Errorf("valor inválido para Status: %s", strValue)
	}
	return nil
}

func (s Status) Value() (driver.Value, error) {
	if !s.IsValid() {
		return nil, errors.New("valor inválido para Status")
	}
	return string(s), nil
}

func (s Status) String() string {
	return string(s)
}

type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status" gorm:"default:'pendente'"`
	Priority    string `json:"priority"`
	ProjectID   uint   `json:"project_id"`
	Project     Project
}
