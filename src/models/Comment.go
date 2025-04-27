package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `json:"content"`
	Solved  bool   `gorm:"default:false"`
	TaskID  uint   `json:"task_id"`
	Task    Task
}
