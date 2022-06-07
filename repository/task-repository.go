package repository

import (
	"time"

	"github.com/zerodev/golang_api/entity"
	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTask(userID uint64) ([]entity.Task, error)
	GetTaskToday(userID uint64) ([]entity.Task, error)
	CreateTask(task entity.Task) (entity.Task, error)
	UpdateTask(task entity.Task) (entity.Task, error)
}

type taskConnection struct {
	connection *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskConnection{
		connection: db,
	}
}

func (db *taskConnection) GetTask(userID uint64) ([]entity.Task, error) {
	var task []entity.Task
	err := db.connection.Where("user_id = ?", userID).Find(&task)
	db.connection.Preload("Label").Find(&task)
	return task, err.Error
}

func (db *taskConnection) GetTaskToday(userID uint64) ([]entity.Task, error) {
	var task []entity.Task
	currentTime := time.Now().Format("2006-01-02")
	err := db.connection.Where("user_id = ? AND datetime LIKE ?", userID, "%"+currentTime+"%").Find(&task).Preload("Label").Find(&task)
	return task, err.Error
}

func (db *taskConnection) CreateTask(task entity.Task) (entity.Task, error) {
	err := db.connection.Save(&task)
	return task, err.Error
}

func (db *taskConnection) UpdateTask(task entity.Task) (entity.Task, error) {
	err := db.connection.Save(&task)
	return task, err.Error
}
