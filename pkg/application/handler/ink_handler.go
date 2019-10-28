package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-hex-sample/pkg/domain/ink"
	"strconv"
)

type InkHandler interface {
	Get(c *gin.Context)
	GetAll(c *gin.Context)
	Add(c *gin.Context)
}

type handler struct {
	service ink.Service
}

func NewInkHandler(s ink.Service) InkHandler {
	return &handler{service: s}
}

func (h *handler) Get(c *gin.Context) {
	queryString := c.Query("id")
	id, err := strconv.ParseInt(queryString, 10, 64)
	if err != nil {
		_ = c.AbortWithError(422, errors.New("id parameter is not an integer"))
		return
	}
	if id == 0 {
		_ = c.AbortWithError(422, errors.New("id query parameter not found"))
		return
	}
	inkResult, err := h.service.GetInk(id)
	if err != nil {
		_ = c.AbortWithError(422, err)
		return
	}
	c.JSON(200, gin.H{"ink": inkResult})
}

func (h *handler) GetAll(c *gin.Context) {
	result, err := h.service.GetAllInks()
	if err != nil {
		_ = c.AbortWithError(500, fmt.Errorf("cound not get inks: %w", err))
		return
	}
	c.JSON(200, gin.H{"inks": result})
}

func (h *handler) Add(c *gin.Context) {
	var inkToAdd ink.Ink
	if err := c.BindJSON(&inkToAdd); err != nil {
		_ = c.AbortWithError(422, fmt.Errorf("could not parse ink JSON: %w", err))
		return
	}
	if err := h.service.AddInk(inkToAdd); err != nil {
		_ = c.AbortWithError(500, fmt.Errorf("error adding ink: %w", err))
		return
	}
	c.Status(201)
}
