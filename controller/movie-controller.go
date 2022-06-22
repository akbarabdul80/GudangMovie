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

type MovieController interface {
	GetMovie(ctx *gin.Context)
	GetMovieByID(ctx *gin.Context)
	CreateMovie(ctx *gin.Context)
	DeleteMovie(ctx *gin.Context)
	WatchMovie(ctx *gin.Context)
}

type movieController struct {
	movieService service.MovieService
	jwtService   service.JWTService
}

func NewMovieController(movieService service.MovieService, jwtService service.JWTService) MovieController {
	return &movieController{
		movieService: movieService,
		jwtService:   jwtService,
	}
}

func (c *movieController) GetMovie(context *gin.Context) {
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

	movie, err := c.movieService.GetMovie(id)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadGateway, response)
		return
	}
	res := helper.BuildResponse(true, "OK!", movie)
	context.JSON(200, res)
}

func (c *movieController) GetMovieByID(context *gin.Context) {
	movie := dto.MovieIDDTO{}
	err := context.ShouldBindJSON(&movie)
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

	res_movie, err := c.movieService.GetMovieByID(id, movie.ID_movie_user)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadGateway, response)
		return
	}
	res := helper.BuildResponse(true, "OK!", res_movie)
	context.JSON(200, res)
}

func (c *movieController) WatchMovie(context *gin.Context) {
	movie := dto.MovieIDDTO{}
	err := context.ShouldBindJSON(&movie)
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

	err_db := c.movieService.WatchMovie(id, movie.ID_movie_user)
	if err_db != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadGateway, response)
		return
	}
	res := helper.BuildResponse(true, "OK!", nil)
	context.JSON(200, res)
}

func (c *movieController) CreateMovie(context *gin.Context) {
	movie := dto.MovieCreateDTO{}
	err := context.ShouldBindJSON(&movie)
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

	movie.UserID = id
	userToCreate, err := c.movieService.CreateMovie(movie)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadGateway, response)
		return
	}
	res := helper.BuildResponse(true, "OK!", userToCreate)
	context.JSON(200, res)
}

func (c *movieController) DeleteMovie(context *gin.Context) {
	movie := dto.MovieIDDTO{}
	err := context.ShouldBindJSON(&movie)
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

	movie.UserID = id
	err_db := c.movieService.DeleteMovie(movie.UserID, movie.ID_movie_user)
	if err_db != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadGateway, response)
		return
	}
	res := helper.BuildResponse(true, "OK!", nil)
	context.JSON(200, res)
}
