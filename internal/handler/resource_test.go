package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"learn_verse/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 🎭 Mock du ResourceService
type MockResourceService struct {
	mock.Mock
}

type mockRealResourceService struct {
	*MockResourceService
}

func (m *mockRealResourceService) Create(ctx context.Context, collectionID *models.ULID, typ string, title string, content, metadata json.RawMessage) (models.Resource, error) {
	return m.MockResourceService.Create(ctx, collectionID, typ, title, content, metadata)
}

func (m *MockResourceService) Create(ctx context.Context, collectionID *models.ULID, typ, title string, content, metadata json.RawMessage) (models.Resource, error) {
	args := m.Called(ctx, collectionID, typ, title, content, metadata)
	return args.Get(0).(models.Resource), args.Error(1)
}

func setupResourceRouter(handler *resourceHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/resources", handler.Create)
	return r
}

func TestResourceHandler_Create(t *testing.T) {
	mockSvc := new(MockResourceService)
	handler := NewResourceHandler(mockSvc)

	collectionID := models.ULID(ulid.Make())
	resourceType := "note"
	title := "Titre de test"
	content := json.RawMessage(`{"body":"Contenu"}`)
	metadata := json.RawMessage(`{"author":"John Doe"}`)

	expected := models.Resource{
		ID:           models.ULID(ulid.Make()),
		CollectionID: &collectionID,
		Type:         resourceType,
		Title:        title,
		Content:      content,
		Metadata:     metadata,
	}

	// Configuration du mock
	mockSvc.On("Create", mock.Anything, &collectionID, resourceType, title, content, metadata).Return(expected, nil)

	// Construction de la requête JSON
	payload := map[string]interface{}{
		"collection_id": collectionID,
		"type":          resourceType,
		"title":         title,
		"content":       content,
		"metadata":      metadata,
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/resources", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router := setupResourceRouter(handler)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockSvc.AssertExpectations(t)
}
