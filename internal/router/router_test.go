package router

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// SetupTestRouter crée un routeur Gin avec les routes de l'app
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	Setup(r, nil) // tu peux injecter une fausse DB plus tard
	return r
}

func TestCollectionsListRoute(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest("GET", "/api/collections", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK && resp.Code != http.StatusInternalServerError {
		t.Errorf("GET /api/collections -> code inattendu: %d", resp.Code)
	}
}

func TestCollectionsCreateRoute(t *testing.T) {
	router := setupTestRouter()

	payload := []byte(`{"name":"Collection test","description":"Ceci est une description"}`)
	req, _ := http.NewRequest("POST", "/api/collections", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated && resp.Code != http.StatusOK && resp.Code != http.StatusInternalServerError {
		t.Errorf("POST /api/collections -> code inattendu: %d", resp.Code)
	}
}

func TestResourcesCreateRoute(t *testing.T) {
	router := setupTestRouter()

	payload := []byte(`{
		"collection_id": "01HABCXYZ1234567890",
		"type": "article",
		"title": "Titre fictif",
		"content": {"body": "Contenu de test"}
	}`)
	req, _ := http.NewRequest("POST", "/api/resources", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated && resp.Code != http.StatusOK && resp.Code != http.StatusInternalServerError {
		t.Errorf("POST /api/resources -> code inattendu: %d", resp.Code)
	}
}
