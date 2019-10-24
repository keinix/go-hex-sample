package main

import (
	"github.com/gin-gonic/gin"
	"go-hex-sample/pkg/data/psql"
	"go-hex-sample/pkg/handler"
	"go-hex-sample/pkg/ink"
	"log"
)

func main() {
	if err := psql.Migrate(); err != nil {
		log.Fatal(err)
	}
	inkRepo := psql.NewInkRepository()
	service := ink.NewService(inkRepo)
	inkHandler := handler.NewInkHandler(service)
	r := gin.Default()
	r.GET("/ink", inkHandler.Get)
	r.POST("/ink", inkHandler.Add)
	r.GET("/inks", inkHandler.GetAll)
	err := r.Run(":8080")
	if err != nil {
		log.Panicf("could not start router %v", err)
	}
}

//tsukushi := ink.Ink{
//Name:        "Tsukushi",
//ColorFamily: ink.Brown,
//}
//if err := service.AddInk(tsukushi); err != nil {
//log.Fatalf("error adding ink: %v", err)
//}
//i, err := service.GetInk(1)
//if err != nil {
//log.Fatal(err)
//}
//log.Printf("Get one: %v", i)
//result, err := service.GetAllInks()
//if err != nil {
//log.Fatal(err)
//}
//log.Printf("Get All: %v", result)
