package main

import (
	"golang-crud/config"
	"golang-crud/controller"
	"golang-crud/helper"
	"golang-crud/model"
	"golang-crud/repository"
	"golang-crud/router"
	"golang-crud/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Started Server!")

	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	tagsRepository := repository.NewTagsRepositoryImpl(db)

	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	tagsController := controller.NewTagsController(tagsService)

	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
