package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/zerodev/golang_api/dto"
	"github.com/zerodev/golang_api/helper"
	"github.com/zerodev/golang_api/service"
)

type UserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) Update(context *gin.Context) {
	user := dto.UserUpdateDTO{}
	err := context.ShouldBindJSON(&user)
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
		panic(err.Error())
	}
	user.ID = id
	userToUpdate, err := c.userService.Update(user)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadGateway, response)
		return
	}
	res := helper.BuildResponse(true, "OK!", userToUpdate)
	context.JSON(200, res)
}

func (c *userController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	user := c.userService.Profile(fmt.Sprintf("%v", claims["user_id"]))
	res := helper.BuildResponse(true, "OK!", user)
	context.JSON(200, res)
}
