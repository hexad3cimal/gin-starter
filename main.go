package main

import (
	"gin-starter/config"
	"gin-starter/controllers"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	uuid "github.com/twinj/uuid"
)

func generateContextId() gin.HandlerFunc {
	return func(c *gin.Context) {
		contextId := uuid.NewV4()
		c.Writer.Header().Set("X-Context-Id", contextId.String())
		c.Next()
	}
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}
		c.Next()
	}
}

func main() {

	config.InitLogger()
	router := gin.Default()
	router.Use(CORS())
	router.Use(generateContextId())
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	config.InitDB()
	v1 := router.Group("/v1")
	{
		user := new(controllers.Api)
		v1.POST("/user/login", user.Login)
		v1.POST("/user/register", user.Register)
		auth := new(controllers.AuthController)
		v1.POST("/token/refresh", auth.Refresh)
	}
	router.Run(":" + config.GetEnvValue("port"))
}
