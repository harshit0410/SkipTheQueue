package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	pkg "skipthequeue/pkg/routers"
	"skipthequeue/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("Starting")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Unable to load .env File", err)
		panic(err.Error())
	}

	initMysqlDB()
	sqlDB, err := utils.DB.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close()
	utils.AutoMigrate()

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "ping"})
	})

	pkg.InitRoutes(router)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Print("Starting server at port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}

func initMysqlDB() {
	config := utils.MysqlConfig{
		Host:     os.Getenv("HOST"),
		User:     os.Getenv("MYSQLUSER"),
		Passowrd: os.Getenv("MYSQLPASSOWRD"),
		DB:       os.Getenv("DATABASE"),
	}

	err := utils.ConnectToMysqlDB(config)

	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
}
