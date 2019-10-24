package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-hex-sample/pkg/ink"
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
		c.JSON(422, gin.H{"error": "id parameter is not an integer"})
	}
	if id == 0 {
		c.JSON(422, gin.H{"error": "id query parameter not found"})
		return
	}
	inkResult, err := h.service.GetInk(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err})
		return
	}
	c.JSON(200, gin.H{"ink": inkResult})
}

func (h *handler) GetAll(c *gin.Context) {
	result, err := h.service.GetAllInks()
	if err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("cound not get inks: %v", err)})
		return
	}
	c.JSON(200, gin.H{"inks": result})
}

func (h *handler) Add(c *gin.Context) {
	var inkToAdd ink.Ink
	if err := c.BindJSON(&inkToAdd); err != nil {
		c.JSON(422, gin.H{"error": fmt.Sprintf("could not parse ink JSON: %v", err)})
		return
	}
	if err := h.service.AddInk(inkToAdd); err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("error adding ink: %v", err)})
		return
	}
	c.Status(201)
}
