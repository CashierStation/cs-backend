package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"csbackend/api"
	"csbackend/db"
	"csbackend/util"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-tty"

	docs "csbackend/docs"

	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file! Skipping...")
	}

	log.Println("Connecting to database...")
	db.Connect()

	doMigration := util.IsFlagPassed("migrate")

	if doMigration {
		log.Println("Migrating database...")
		db.Migrate()
	}

	database, err := db.DB.DB()
	if err != nil {
		panic("failed to connect database")
	}

	var owner db.Owner
	db.DB.First(&owner)

	println(owner.Email)

	defer database.Close()

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/"
	apiGroup := r.Group("/api/")
	{
		apiGroup.GET("/ping", api.HelloWorld)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	mode := os.Getenv("GIN_MODE")
	if mode != "release" {
		go check_termination()
	}

	r.Run()
}

func check_termination() {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	for {
		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		if r == 'q' {
			fmt.Println("Terminating...")
			os.Exit(0)
			break
		}
	}
}
