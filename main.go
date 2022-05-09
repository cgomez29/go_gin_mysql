package main

import (
	"github.com/cgomez29/api-gin/config"
	"github.com/cgomez29/api-gin/controller"
	"github.com/cgomez29/api-gin/middleware"
	"github.com/cgomez29/api-gin/repository"
	"github.com/cgomez29/api-gin/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	userService    service.UserService       = service.NewUserService(userRepository)
	authService    service.AuthService       = service.NewAuthService(userRepository)
	autController  controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
)

func main() {

	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	auth_routes := r.Group("api/auth")
	{
		auth_routes.POST("/login", autController.Login)
		auth_routes.POST("/register", autController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	r.Run(":8085")
}
