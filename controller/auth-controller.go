package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zerodev/golang_api/dto"
	"github.com/zerodev/golang_api/entity"
	"github.com/zerodev/golang_api/helper"
	"github.com/zerodev/golang_api/service"
)

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	RefreshToken(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	err := ctx.ShouldBind(&loginDTO)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusOK, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)

	if v, ok := authResult.(entity.User); ok {
		token, rt := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = token
		v.RefreshToken = rt
		response := helper.BuildResponse(true, "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helper.BuildErrorResponse("Please check again yout credetial", "Invalid credetial", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	err := ctx.ShouldBind(&registerDTO)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to process request", err.Error(), helper.EmptyObj{})
		ctx.JSON(http.StatusOK, response)
		return
	}

	if c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helper.BuildErrorResponse("Failed to process request", "Duplicate Email", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusConflict, response)
		return
	} else {
		createUser := c.authService.RegisterUser(registerDTO)
		token, rt := c.jwtService.GenerateToken(strconv.FormatUint(createUser.ID, 10))
		createUser.Token = token
		createUser.RefreshToken = rt
		response := helper.BuildResponse(true, "Ok!", createUser)
		ctx.JSON(http.StatusCreated, response)
		return
	}
}

func (c *authController) RefreshToken(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	authRefreshHeader := ctx.GetHeader("Authorization_Refresh")

	token, errToken := c.jwtService.ValidateToken(authHeader)
	tokenRefresh, errTokenRefersh := c.jwtService.ValidateRefreshToken(authRefreshHeader)

	if errToken.Error() != "Token is expired" || errTokenRefersh != nil {
		res := helper.BuildErrorResponse("Failed to process request", errToken.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}

	success, userID := c.jwtService.ValidatePlayload(*token, *tokenRefresh)

	if !success {
		res := helper.BuildErrorResponse("Failed to process request2", "Invalid token", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	} else {
		result := c.authService.FindByID(userID)
		if v, ok := result.(entity.User); ok {
			token, rt := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
			v.Token = token
			v.RefreshToken = rt
			response := helper.BuildResponse(true, "OK!", v)
			ctx.JSON(http.StatusOK, response)
			return
		}
	}
}
