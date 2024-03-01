package api

import (
	_ "go-exam/api-gateway/api/docs" // swag
	v1 "go-exam/api-gateway/api/handlers/v1"
	"go-exam/api-gateway/config"
	"go-exam/api-gateway/pkg/logger"
	"go-exam/api-gateway/services"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

// @Title Welcome to ProductAPI
// @Version 1.0
// @Description This is a example of Social Network
// @Host localhost:8080
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	api := router.Group("/v1")
	// users
	api.POST("/users", handlerV1.CreateUser)
	api.GET("/users/:id", handlerV1.GetUser)
	api.GET("/users", handlerV1.GetAll)
	api.PUT("/users/:id", handlerV1.UpdateUser)
	api.DELETE("/users/:id", handlerV1.DeleteUser)

	// product
	api.POST("/create", handlerV1.Create)
	api.GET("/get/:id", handlerV1.Get)
	api.GET("/all", handlerV1.GetAll)
	api.PUT("/update/:id", handlerV1.Update)
	api.DELETE("/delete/:id", handlerV1.Delete)

	// register
	api.POST("/sign", handlerV1.SignUp)
	api.GET("/login", handlerV1.LogIn)
	api.GET("/verification", handlerV1.Verification)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
