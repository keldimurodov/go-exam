package v1

import (
	"context"
	"fmt"
	"go-exam/api-gateway/api/handlers/models"
	pbu "go-exam/api-gateway/genproto/user"
	l "go-exam/api-gateway/pkg/logger"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// Sign
// @Summary Sign User
// @Security ApiKeyAuth
// @Description Sign - Api for registring users
// @Tags registr
// @Accept json
// @Produce json
// @Param registr body models.UserDetail true "UserDetail"
// @Success 200 {object} models.ResponseUser
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/sign/ [post]
func (h *handlerV1) SignUp(c *gin.Context) {
	var (
		body        models.UserDetail
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	fmt.Println(body)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().Sign(ctx, &pbu.UserDetail{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  body.Password,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)

}

// LogIn
// @Summary LogIn User
// @Security ApiKeyAuth
// @Description LogIn - Api for login users
// @Tags registr
// @Accept json
// @Produce json
// @Param email query string true "Email"
// @Param password query string true "Password"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/login [get]
func (h *handlerV1) LogIn(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	email := c.Query("email")
	password := c.Query("password")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	responseUser, err := h.serviceManager.UserService().Login(ctx, &pbu.LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user info", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, responseUser)
}

// Verification
// @Summary Verification User
// @Security ApiKeyAuth
// @Description LogIn - Api for verification users
// @Tags registr
// @Accept json
// @Produce json
// @Param email query string true "Email"
// @Param code query string true "Code"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/verification [get]
func (h *handlerV1) Verification(c *gin.Context) {

	email := c.Query("email")
	code := c.Query("code")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	responseUser, err := h.serviceManager.UserService().Verification(ctx, &pbu.VerificationUserRequest{
		Email: email,
		Code:  code,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user info", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, responseUser)
}
