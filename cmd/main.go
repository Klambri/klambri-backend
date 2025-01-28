package main

import (
	"fmt"
	"log"

	"github.com/Klambri/klambri-backend/internal/configuration"
	"github.com/Klambri/klambri-backend/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := configuration.ReadConfig("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	v1 := r.Group("/api/v1")

	routes.SetupRoutes(v1)

	if err := r.Run(fmt.Sprintf(":%v", config.Server.Port)); err != nil {
		panic(err)
	}
}
