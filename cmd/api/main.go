package main

import (
	"go-hex-sample/pkg/data/psql"
	"go-hex-sample/pkg/ink"
	"log"
)

func main() {
	if err := psql.Migrate(); err != nil {
		log.Fatal(err)
	}
	inkRepo := psql.NewInkRepository()
	service := ink.NewService(inkRepo)
	tsukushi := ink.Ink{
		Name:        "Tsukushi",
		ColorFamily: ink.Brown,
	}
	if err := service.AddInk(tsukushi); err != nil {
		log.Fatalf("error adding ink: %v", err)
	}
	i, err := service.GetInk(1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Get one: %v", i)
	result, err := service.GetAllInks()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Get All: %v", result)
}
