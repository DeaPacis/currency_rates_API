package main

import (
	"app/internal/config"
	"app/internal/handler"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	router := gin.Default()

	h := handler.NewHandler(cfg)

	router.GET("/info", h.Info)
	router.GET("/info/currency", h.Currency)

	err := router.Run(":" + cfg.Port)
	if err != nil {
		log.Printf("Error starting service: %s", err)
		return
	} else {
		log.Println("Service started successfully")
	}
}
