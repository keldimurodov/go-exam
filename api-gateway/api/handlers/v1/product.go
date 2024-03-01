package v1

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	models "go-exam/api-gateway/api/handlers/models"
	"go-exam/api-gateway/api/handlers/tokens"
	pbp "go-exam/api-gateway/genproto/product"
	l "go-exam/api-gateway/pkg/logger"
	"go-exam/api-gateway/pkg/utils"
)

// CreateProduct ...
// @Summary CreateProduct ...
// @Security ApiKeyAuth
// @Description Api for creating a new product
// @Tags product
// @Accept json
// @Produce json
// @Param Product body models.Product true "createProduct"
// @Success 200 {object} models.Product
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/create/ [post]
func (h *handlerV1) Create(c *gin.Context) {
	var (
		body        models.Product
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	h.jwthandler = tokens.JWTHandler{
		Sub:       body.ProductName,
		Role:      "admin",
		SigninKey: "admin",
		Log:       h.log,
	}

	access, refresh, err := h.jwthandler.GenerateAuthJWT()
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "error while generating jwt",
		})
		h.log.Error("error generate new jwt tokens", l.Error(err))
		return
	}
	fmt.Println("Tokens are working well")

	response, err := h.serviceManager.ProductService().Create(ctx, &pbp.Product{
		ProductName:  body.ProductName,
		ProductPrice: body.ProductPrice,
		ProductAbout: body.ProductAbout,
		RefreshToken: refresh,
	})

	respBody := &models.CreateProductResponse{
		Id:           response.Id,
		ProductName:  response.ProductName,
		ProductPrice: response.ProductPrice,
		ProductAbout: response.ProductAbout,
		CreatedAt:    response.CreatedAt,
		RefreshToken: response.RefreshToken,
		AccesToken:   access,
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, respBody)
}

// Get get product by id
// @Summary GetProduct
// @Security ApiKeyAuth
// @Description Api for getting product by id
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Product
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/get/{id} [get]
func (h *handlerV1) Get(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.ProductService().Get(
		ctx, &pbp.GetRequest{
			Id: id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// ListProduct returns list of products
// @Summary All products
// @Security ApiKeyAuth
// @Description Api returns list of products
// @Tags product
// @Accept json
// @Produce json
// @Param page path int64 true "Page"
// @Param limit path int64 true "Limit"
// @Succes 200 {object} models.Product
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/all/ [get]
func (h *handlerV1) GetAll(c *gin.Context) {
	queryParams := c.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("failed to parse query params json" + errStr[0])
		return
	}

	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.ProductService().GetAll(
		ctx, &pbp.GetAllRequest{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list products", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateProduct updates product by id
// @Summary UpdateProduct
// @Security ApiKeyAuth
// @Description Api returns updates user
// @Tags product
// @Accept json
// @Produce json
// @Param Product body models.Product true "UpdateProduct"
// @Succes 200 {Object} models.Product
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/update/{id} [put]
func (h *handlerV1) Update(c *gin.Context) {
	var (
		body        pbp.Product
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
	body.Id = c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.ProductService().Update(ctx, &pbp.Product{
		Id:           body.Id,
		ProductName:  body.ProductName,
		ProductPrice: body.ProductPrice,
		ProductAbout: body.ProductAbout,
		CreatedAt:    body.CreatedAt,
		UpdetedAt:    body.UpdetedAt,
		DeletedAt:    body.DeletedAt,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteProduct deletes product by id
// @Summary DeleteProduct
// @Security ApiKeyAuth
// @Description Api deletes product
// @Tags product
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Succes 200 {Object} model.Product
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/delete/{id} [delete]
func (h *handlerV1) Delete(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.ProductService().Delete(
		ctx, &pbp.GetRequest{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
