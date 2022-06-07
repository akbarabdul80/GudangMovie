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

type LabelController interface {
	GetLabel(ctx *gin.Context)
	CreateLabel(ctx *gin.Context)
	UpdateLabel(ctx *gin.Context)
}

type labelController struct {
	labelService service.LabelService
	jwtService   service.JWTService
}

func NewLabelController(labelService service.LabelService, jwtService service.JWTService) LabelController {
	return &labelController{
		labelService: labelService,
		jwtService:   jwtService,
	}
}

func (c *labelController) GetLabel(context *gin.Context) {
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

	label, err := c.labelService.GetLabel(id)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadGateway, response)
		return
	}
	res := helper.BuildResponse(true, "OK!", label)
	context.JSON(200, res)
}

func (c *labelController) CreateLabel(context *gin.Context) {
	label := dto.LabelCreateDTO{}
	err := context.ShouldBindJSON(&label)
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

	label.UserID = id
	userToCreate, err := c.labelService.CreateLabel(label)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadGateway, response)
		return
	}
	res := helper.BuildResponse(true, "OK!", userToCreate)
	context.JSON(200, res)
}

func (c *labelController) UpdateLabel(context *gin.Context) {
	label := dto.LabelUpdateDTO{}
	err := context.ShouldBindJSON(&label)
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

	label.UserID = id
	userToUpdate, err := c.labelService.UpdateLabel(label)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadGateway, response)
		return
	}
	res := helper.BuildResponse(true, "OK!", userToUpdate)
	context.JSON(200, res)
}
