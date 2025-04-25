package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"learn_verse/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// MockCollectionService simule le comportement du vrai service
type MockCollectionService struct {
	mock.Mock
}

func (m *MockCollectionService) Update(ctx context.Context, collection models.ResourceCollection) (models.ResourceCollection, error) {
	args := m.Called(ctx, collection)
	return args.Get(0).(models.ResourceCollection), args.Error(1)
}

func (m *MockCollectionService) Delete(ctx context.Context, id models.ULID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockCollectionService) Create(ctx context.Context, coll models.ResourceCollection) (models.ResourceCollection, error) {
	args := m.Called(ctx, coll)
	return args.Get(0).(models.ResourceCollection), args.Error(1)
}

func (m *MockCollectionService) Get(ctx context.Context, id models.ULID) (models.ResourceCollection, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(models.ResourceCollection), args.Error(1)
}

func (m *MockCollectionService) List(ctx context.Context) ([]models.ResourceCollection, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.ResourceCollection), args.Error(1)
}

func setupRouterWithHandler(h *collectionHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/collections", h.Create)
	r.GET("/collections", h.List)
	r.GET("/collections/:id", h.Get)
	return r
}

func TestCollectionHandler_Create(t *testing.T) {
	mockSvc := new(MockCollectionService)
	handler := NewCollectionHandler(mockSvc)

	description := "Test Collection"
	input := models.ResourceCollection{
		Name:        "Test Collection",
		Description: &description,
	}

	expected := input
	expected.ID = models.ULID(ulid.Make())
	expected.CreatedAt = time.Now()
	expected.UpdatedAt = expected.CreatedAt

	mockSvc.On("Create", mock.Anything, input).Return(expected, nil)

	body, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/collections", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router := setupRouterWithHandler(handler)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestCollectionHandler_Get(t *testing.T) {
	mockSvc := new(MockCollectionService)
	handler := NewCollectionHandler(mockSvc)

	id := models.ULID(ulid.Make())
	description := "Test Collection"
	expected := models.ResourceCollection{
		ID:          id,
		Name:        "Mocked",
		Description: &description,
	}

	mockSvc.On("Get", mock.Anything, id).Return(expected, nil)
	modelULID := models.ULID(id)
	strValue, err := modelULID.Value()
	req, err := http.NewRequest("GET", fmt.Sprintf("/collections/%s", strValue), nil)
	fmt.Print(err)
	w := httptest.NewRecorder()
	router := setupRouterWithHandler(handler)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)
}

func TestCollectionHandler_List(t *testing.T) {
	mockSvc := new(MockCollectionService)
	handler := NewCollectionHandler(mockSvc)

	desc1 := "Desc 1"
	desc2 := "Desc 2"
	list := []models.ResourceCollection{
		{
			ID:          models.ULID(ulid.Make()),
			Name:        "Item 1",
			Description: &desc1,
		},
		{
			ID:          models.ULID(ulid.Make()),
			Name:        "Item 2",
			Description: &desc2,
		},
	}

	mockSvc.On("List", mock.Anything).Return(list, nil)

	req, _ := http.NewRequest("GET", "/collections", nil)
	w := httptest.NewRecorder()
	router := setupRouterWithHandler(handler)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockSvc.AssertExpectations(t)
}
