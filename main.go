package main

import (
	"github.com/gin-gonic/gin"
	"my-first-golang-program/internal/adapters/input"
	"my-first-golang-program/internal/adapters/output"
	"my-first-golang-program/internal/application"
)

func main() {

	cdrRepository := output.NewCdrInMemoryRepository()
	getCdrUseCase := application.NewGetCdrUseCase(cdrRepository)
	getCdrHttpHandler := input.NewHTTPHandler(getCdrUseCase)

	router := gin.New()
	router.GET("/cdr/:id", getCdrHttpHandler.Get)

	router.Run(":8080")
}
