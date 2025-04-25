package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"learn_verse/internal/models"
	"learn_verse/internal/service"
	"net/http"
)

// func NewResourceHandler(svc *service.ResourceService) *resourceHandler { return &resourceHandler{svc: svc} }

//type resourceHandler struct { svc *service.ResourceService }

func NewResourceHandler(svc service.ResourceServiceInterface) *resourceHandler {
	return &resourceHandler{svc: svc}
}

type resourceHandler struct {
	svc service.ResourceServiceInterface
}

func (h *resourceHandler) Create(c *gin.Context) {
	var in struct {
		CollectionID *models.ULID    `json:"collection_id"`
		Type         string          `json:"type" binding:"required"`
		Title        string          `json:"title" binding:"required"`
		Content      json.RawMessage `json:"content" binding:"required"`
		Metadata     json.RawMessage `json:"metadata"`
	}
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.svc.Create(c.Request.Context(), in.CollectionID, in.Type, in.Title, in.Content, in.Metadata)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}
