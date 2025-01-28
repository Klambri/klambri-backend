package main

import (
	"fmt"
	"log"

	"github.com/Klambri/klambri-backend/internal/configuration"
	"github.com/Klambri/klambri-backend/internal/routes"
	"github.com/gin-gonic/gin"

	docs "github.com/Klambri/klambri-backend/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	config, err := configuration.ReadConfig("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := r.Group("/api/v1")
	{
		//LDAP routes
		routes.User(v1.Group("/user"))

		//Ansible routes
		routes.Playbook(v1.Group("playbook"))
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	if err := r.Run(fmt.Sprintf(":%v", config.Server.Port)); err != nil {
		panic(err)
	}
}
