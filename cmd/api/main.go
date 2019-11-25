package main

import (
	"github.com/gin-gonic/gin"
	"go-hex-sample/pkg/application/handler"
	"go-hex-sample/pkg/application/middlewear"
	"go-hex-sample/pkg/data/psql"
	"go-hex-sample/pkg/data/redis"
	"go-hex-sample/pkg/domain/ink"
	"go-hex-sample/pkg/domain/login"
	"log"
)

func main() {
	if err := psql.Migrate(); err != nil {
		log.Fatal(err)
	}
	inkRepo := psql.NewInkRepository()
	inkService := ink.NewService(inkRepo)
	inkHandler := handler.NewInkHandler(inkService)

	loginRepo := psql.NewUserRepository()
	tokenCache := redis.NewTokenCache()
	loginService := login.NewService(loginRepo, tokenCache)
	loginHandler := handler.NewLoginHandler(loginService)

	r := gin.Default()
	r.Use(middlewear.HandleError())
	r.POST("/login", loginHandler.Login)

	basicAuth := r.Group("/")
	basicAuth.Use(middlewear.AuthRequired(tokenCache))
	{
		basicAuth.GET("/ink", inkHandler.Get)
		basicAuth.POST("/ink", inkHandler.Add)
		basicAuth.GET("/inks", inkHandler.GetAll)
	}

	err := r.Run(":8080")
	if err != nil {
		log.Panicf("could not start router %v", err)
	}
}
