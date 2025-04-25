package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"learn_verse/internal/models"
	"learn_verse/internal/service"
	"net/http"
	"strings"
)

//func NewCollectionHandler(svc *service.CollectionService) *collectionHandler { return &collectionHandler{svc: svc} }

//type collectionHandler struct { svc *service.CollectionService }

func NewCollectionHandler(svc service.CollectionServiceInterface) *collectionHandler {
	return &collectionHandler{svc: svc}
}

type collectionHandler struct {
	svc service.CollectionServiceInterface
}

func (h *collectionHandler) Create(c *gin.Context) {
	in := models.ResourceCollection{}
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	coll, err := h.svc.Create(c.Request.Context(), in)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, coll)
}
func (h *collectionHandler) Get(c *gin.Context) {
	idParam := c.Param("id")
	id, err := ulid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}
	coll, err := h.svc.Get(c.Request.Context(), models.ULID(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Collection non trouvée"})
		return
	}
	c.JSON(http.StatusOK, coll)
}

// List GET /collections
func (h *collectionHandler) List(c *gin.Context) {
	list, err := h.svc.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}

// DELETE /collection/:id

func (h *collectionHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := ulid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}
	err = h.svc.Delete(c.Request.Context(), models.ULID(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Collection non trouvée"})
		return
	}
	c.JSON(http.StatusNoContent, "Deleted")
}

func (h *collectionHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := ulid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}
	var in models.ResourceCollection
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	in.ID = models.ULID(id)
	updated, err := h.svc.Update(c.Request.Context(), in)
	if err != nil {
		if strings.Contains(err.Error(), "collection not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": "Collection non trouvée ou déjà supprimée"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}
