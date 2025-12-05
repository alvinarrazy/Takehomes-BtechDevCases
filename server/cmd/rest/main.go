package main

import (
	"log"
	"server/middleware"
	"server/route"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.SetFlags(log.Lshortfile)

	router := gin.Default()
	router.Use(middleware.Cors())

	router.Static("/docs", "./dist")

	route.Routes(router)

	log.Fatal(router.Run(":8080"))
}
