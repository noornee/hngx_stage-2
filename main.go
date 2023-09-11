package main

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/noornee/hngx_stage-2/internal/config"
	"github.com/noornee/hngx_stage-2/pkg/repository/storage/mongodb"
	"github.com/noornee/hngx_stage-2/pkg/router"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = config.Setup().PORT
	}

	db, _ := mongodb.ConnectToDB()
	validate := validator.New()

	r := router.Setup(db, validate)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err.Error())
	}
}
