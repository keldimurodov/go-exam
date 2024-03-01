package v1

import (
	models "go-exam/api-gateway/api/handlers/models"
	tokens "go-exam/api-gateway/api/handlers/tokens"
	"go-exam/api-gateway/config"
	"go-exam/api-gateway/pkg/logger"
	"go-exam/api-gateway/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlerV1 struct {
	log            logger.Logger
	serviceManager services.IServiceManager
	cfg            config.Config
	jwthandler     tokens.JWTHandler
}

// handlerV1Config ...
type HandlerV1Config struct {
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Cfg            config.Config
	JWTHandler     tokens.JWTHandler
}

// New ...
func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:            c.Logger,
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
		jwthandler:     c.JWTHandler,
	}
}

func handleBadRequestWithErrorMessage(c *gin.Context, l logger.Logger, err error, message string) bool {
	if err != nil {
		c.JSON(http.StatusBadRequest, models.StandardErrorModel{
			Error: models.Error{
				Message: "Incorrect data supplied",
			},
		})
		l.Error(message, logger.Error(err))
		return true
	}
	return false
}
