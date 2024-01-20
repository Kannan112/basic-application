package http

import (
	"github.com/gin-gonic/gin"
	_ "github.com/kannan112/token-application/cmd/api/docs"
	"github.com/kannan112/token-application/pkg/api/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

// @title Clinic Application API
// @title Clean- Application API
// @version 1.0.0
// @description  Backend API Golang using Clean Code architecture
// @contact.name API Support
// @contact.email				abhinandarun11@gmail.com
// @BasePath /
// @query.collection.format multi
func NewServerHTTP(userHandler *handler.UserHandler) *ServerHTTP {
	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())

	// Swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Auth middleware
	api := engine.Group("/api")

	api.POST("create-ticket/:id", userHandler.CreateTicket)
	api.GET("get-ticket/:id", userHandler.ShowTicket)
	api.POST("clinic-create", userHandler.CreateClinic)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":8080")
}
