package service

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/zerodev/golang_api/dto"
	"github.com/zerodev/golang_api/entity"
	"github.com/zerodev/golang_api/repository"
)

type TaskService interface {
	GetTask(userID uint64) ([]entity.Task, error)
	GetTaskToday(userID uint64) ([]entity.Task, error)
	CreateTask(task dto.TaskCreateDTO) (entity.Task, error)
	UpdateTask(task dto.TaskUpdateDTO) (entity.Task, error)
}

type taskService struct {
	taskRepository repository.TaskRepository
}

func NewTaskService(taskRepo repository.TaskRepository) TaskService {
	return &taskService{
		taskRepository: taskRepo,
	}
}

func (service *taskService) GetTask(userID uint64) ([]entity.Task, error) {
	res, err := service.taskRepository.GetTask(userID)
	return res, err
}

func (service *taskService) GetTaskToday(userID uint64) ([]entity.Task, error) {
	res, err := service.taskRepository.GetTaskToday(userID)
	return res, err
}

func (service *taskService) CreateTask(task dto.TaskCreateDTO) (entity.Task, error) {
	taskToCreate := entity.Task{}
	err := smapping.FillStruct(&taskToCreate, smapping.MapFields((&task)))
	if err != nil {
		log.Fatalf("Failed mapping %v", err.Error())
	}

	res, err := service.taskRepository.CreateTask(taskToCreate)

	return res, err
}

func (service *taskService) UpdateTask(task dto.TaskUpdateDTO) (entity.Task, error) {
	taskToUpdate := entity.Task{}
	err := smapping.FillStruct(&taskToUpdate, smapping.MapFields((&task)))
	if err != nil {
		log.Fatalf("Failed mapping %v", err.Error())
	}

	res, err := service.taskRepository.UpdateTask(taskToUpdate)

	return res, err
}
