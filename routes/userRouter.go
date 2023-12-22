package routes

import(
	controller "go-jwt-project/controllers"
	"github.com/gin-gonic/gin"
	"go-jwt-project/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users",controller.GetUsers())
	incomingRoutes.GET("/users/:user_id",controller.GetUser())
}

