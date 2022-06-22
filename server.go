package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zerodev/golang_api/config"
	"github.com/zerodev/golang_api/controller"
	"github.com/zerodev/golang_api/middleware"
	"github.com/zerodev/golang_api/repository"
	"github.com/zerodev/golang_api/service"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
	// Respository
	userRepository  repository.UserRepository  = repository.NewUserRepository(db)
	movieRepository repository.MovieRepository = repository.NewMovieRepository(db)

	// Service
	jwtService   service.JWTService   = service.NewJWTService()
	authService  service.AuthService  = service.NewAuthService(userRepository)
	userService  service.UserService  = service.NewUserService(userRepository)
	movieService service.MovieService = service.NewMovieService(movieRepository)

	// Controller
	authController  controller.AuthController  = controller.NewAuthController(authService, jwtService)
	userController  controller.UserController  = controller.NewUserController(userService, jwtService)
	movieController controller.MovieController = controller.NewMovieController(movieService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
		authRoutes.GET("/refresh-token", authController.RefreshToken)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.POST("/update", userController.Update)

		// Movie
		userRoutes.GET("/movie", movieController.GetMovie)
		userRoutes.GET("/movie-id", movieController.GetMovieByID)
		userRoutes.PUT("/movie", movieController.CreateMovie)
		userRoutes.DELETE("/movie", movieController.DeleteMovie)
		userRoutes.PATCH("/movie	", movieController.WatchMovie)
	}

	r.Run("0.0.0.0:8081")
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
