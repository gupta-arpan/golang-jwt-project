package routes

import(
	"github.com/gin-gonic/gin";
	controllers "github.com/gupta-arpan/golang-jwt-project/controllers";
	middlewares "github.com/gupta-arpan/golang-jwt-project/middlewares";
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middlewares.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}