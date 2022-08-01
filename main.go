package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"transactions/configs"
	"transactions/controllers"
)

func main() {
	viper.SetConfigFile("dev.env")
	viper.ReadInConfig()
	configs.CreateTables()
	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/transactions", controllers.GetTransactions)
	router.Run(":80")
}
