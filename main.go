package main

import (
	"fmt"
	"os"

	"github.com/tugas-workshop/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	doRequest := new(controllers.RequestController)

	if viper.GetString("mode") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	router.GET("/", func(c *gin.Context) {
		name, err := os.Hostname()
		if err != nil {
			panic(err)
		}
		c.JSON(200, gin.H{
			"message":   "TV Show API",
			"served_by": name,
		})
	})
	v1 := router.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {

			c.JSON(200, gin.H{
				"message": "Welcome to my API",
			})
		})
		v1.GET("/schedule", doRequest.Schedule)
		v1.GET("/search", doRequest.Search)
	}

	router.Run(":" + viper.GetString("port"))
}
