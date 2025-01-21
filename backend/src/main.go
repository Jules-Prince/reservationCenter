package main

import (
	"hexagone/reservationCenter/src/routes"
	"hexagone/reservationCenter/src/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
  utils.InitRedis()

	router := gin.Default()

	router.Use(cors.Default())

	// Routes pour les disponibilités
	router.GET("/availabilities", routes.GetAvailabilities)
	router.POST("/availabilities", routes.CreateAvailability)

	// Démarrer le serveur
	router.Run(":8080")
}
