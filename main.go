package main

import (
	//routes "github.com/gupta-arpan/golang-jwt-project/routes";
	"os";
	"github.com/gin-gonic/gin";
)

func main(){
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	//routes.AuthRoutes(router)
	//routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted to API 1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted to API 2"})
	})

	router.Run(":" + port)
}