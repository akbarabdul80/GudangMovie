package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/zerodev/golang_api/dto"
	"github.com/zerodev/golang_api/helper"
	"github.com/zerodev/golang_api/service"
)

type TaskController interface {
	GetTask(ctx *gin.Context)
	GetTaskToday(ctx *gin.Context)
	CreateTask(ctx *gin.Context)
	UpdateTask(ctx *gin.Context)
}

type taskController struct {
	taskService service.TaskService
	jwtService  service.JWTService
}

func NewTaskController(taskService service.TaskService, jwtService service.JWTService) TaskController {
	return &taskController{
		taskService: taskService,
		jwtService:  jwtService,
	}
}

func (c *taskController) GetTask(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(claims["user_id"].(string), 10, 64)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadGateway, response)
		return
	}

	task, err := c.taskService.GetTask(id)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadGateway, response)
		return
	}
	res := helper.BuildResponse(true, "OK!", task)
	context.JSON(200, res)
}

func (c *taskController) GetTaskToday(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(claims["user_id"].(string), 10, 64)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadGateway, response)
		return
	}

	task, err := c.taskService.GetTaskToday(id)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadGateway, response)
		return
	}
	res := helper.BuildResponse(true, "OK!", task)
	context.JSON(200, res)
}

func (c *taskController) CreateTask(context *gin.Context) {
	task := dto.TaskCreateDTO{}
	err := context.ShouldBindJSON(&task)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadGateway, response)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(claims["user_id"].(string), 10, 64)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadGateway, response)
		return
	}

	task.UserID = id
	taskCreate, err := c.taskService.CreateTask(task)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadGateway, response)
		return
	}
	res := helper.BuildResponse(true, "OK!", taskCreate)
	context.JSON(200, res)
}

func (c *taskController) UpdateTask(context *gin.Context) {
	task := dto.TaskUpdateDTO{}
	err := context.ShouldBindJSON(&task)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadGateway, response)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		response := helper.BuildErrorResponse("Failed to process request", errToken.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadGateway, response)
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(claims["user_id"].(string), 10, 64)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadGateway, response)
		return
	}

	task.UserID = id
	userToUpdate, err := c.taskService.UpdateTask(task)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadGateway, response)
		return
	}

	res := helper.BuildResponse(true, "OK!", userToUpdate)
	context.JSON(200, res)
}
