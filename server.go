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
	labelRepository repository.LabelRepository = repository.NewLabelRepository(db)
	taskRepository  repository.TaskRepository  = repository.NewTaskRepository(db)

	// Service
	jwtService   service.JWTService   = service.NewJWTService()
	authService  service.AuthService  = service.NewAuthService(userRepository)
	userService  service.UserService  = service.NewUserService(userRepository)
	labelService service.LabelService = service.NewLabelService(labelRepository)
	taskService  service.TaskService  = service.NewTaskService(taskRepository)

	// Controller
	authController  controller.AuthController  = controller.NewAuthController(authService, jwtService)
	userController  controller.UserController  = controller.NewUserController(userService, jwtService)
	labelController controller.LabelController = controller.NewLabelController(labelService, jwtService)
	taskController  controller.TaskController  = controller.NewTaskController(taskService, jwtService)
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

		// label
		userRoutes.GET("/label", labelController.GetLabel)
		userRoutes.PUT("/label", labelController.CreateLabel)
		userRoutes.PATCH("/label", labelController.UpdateLabel)

		// Task
		userRoutes.GET("/task", taskController.GetTask)
		userRoutes.PUT("/task", taskController.CreateTask)
		userRoutes.PATCH("/task", taskController.UpdateTask)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
