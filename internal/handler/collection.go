package handler

import (
	"github.com/gin-gonic/gin"
	"learn_verse/internal/service"
	"net/http"
)

func NewCollectionHandler(svc *service.CollectionService) *collectionHandler {
	return &collectionHandler{svc: svc}
}

type collectionHandler struct {
	svc *service.CollectionService
}

func (h *collectionHandler) Create(c *gin.Context) {
	var in struct {
		Name        string  `json:"name"`
		Description *string `json:"description"`
	}
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	coll, err := h.svc.Create(c.Request.Context(), in.Name, in.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, coll)
}
