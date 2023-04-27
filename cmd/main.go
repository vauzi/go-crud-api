package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vauzi/go-crud-api/pkg/common/config"
	"github.com/vauzi/go-crud-api/pkg/common/database"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	// viper.SetConfigFile("./pkg/common/envs/.env")
	// viper.ReadInConfig()
	// add env variables as needed
	port := cfg.Port

	dbHost := cfg.DBHost
	dbPort := cfg.DBPort
	dbDatabase := cfg.DBDatabase
	dbUsername := cfg.DBUsername
	dbPassword := cfg.DBPassword

	DBUrl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbDatabase, dbPassword)

	database.Init(DBUrl)

	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"port":  port,
			"dbUrl": "dbUrl",
		})
	})

	router.Run(":3000")
}
