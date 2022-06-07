package dto

import (
	"time"
)

type TaskUpdateDTO struct {
	ID_task   uint64    `json:"id_task" form:"id_task" binding:"required"`
	Name_task string    `json:"name_task" form:"name_task" binding:"required"`
	Datetime  time.Time `json:"datetime" form:"datetime" binding:"required"`
	Note      string    `json:"note" form:"note" binding:"required"`
	Done      *bool     `json:"done" form:"done" binding:"required"`
	UserID    uint64    `json:"user_id,omitempty" form:"user_id,omitempty"`
	LabelID   uint64    `json:"label_id,omitempty" form:"label_id,omitempty"`
}

type TaskCreateDTO struct {
	Name_task string    `json:"name_task" form:"name_task" binding:"required"`
	Datetime  time.Time `json:"datetime" form:"datetime" binding:"required"`
	Note      string    `json:"note" form:"note" binding:"required"`
	UserID    uint64    `json:"user_id,omitempty" form:"user_id,omitempty"`
	LabelID   uint64    `json:"label_id,omitempty" form:"label_id,omitempty"`
}
