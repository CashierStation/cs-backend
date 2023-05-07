package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"csbackend/api"
	"csbackend/db"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-tty"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file! Skipping...")
	}

	log.Println("Connecting to database...")
	db.Connect()

	log.Println("Migrating database...")
	db.Migrate()

	database, err := db.DB.DB()
	if err != nil {
		panic("failed to connect database")
	}

	db.DB.Raw("SELECT 1").Scan(&database)

	defer database.Close()

	r := gin.Default()
	r.GET("/ping", api.HelloWorld)

	go check_termination()

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
