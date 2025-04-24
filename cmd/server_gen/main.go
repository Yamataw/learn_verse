package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"learn_verse/server"
	"log"
	"net/http"
)

// apiServer implémente server.ServerInterface et gère la connexion DB
type apiServer struct {
	db *sql.DB
}

// NewAPIServer construit un serveur API avec pool DB
func NewAPIServer(db *sql.DB) server.ServerInterface {
	return &apiServer{db: db}
}

func (s *apiServer) GetCollections(ctx echo.Context) error {
	// TODO: utiliser s.db pour récupérer les collections
	return ctx.NoContent(http.StatusNotImplemented)
}

func (s *apiServer) PostCollections(ctx echo.Context) error {
	// TODO: utiliser s.db pour créer une collection
	return ctx.NoContent(http.StatusNotImplemented)
}

func (s *apiServer) GetCollectionsId(ctx echo.Context, id string) error {
	// TODO: récupérer une collection par ID en DB
	return ctx.NoContent(http.StatusNotImplemented)
}

func (s *apiServer) PostResources(ctx echo.Context) error {
	// TODO: insérer une ressource en DB
	return ctx.NoContent(http.StatusNotImplemented)
}

func main() {
	user := "postgres"
	password := "pwd"
	host := "localhost"
	port := "5432"
	dbname := "learn_verse"
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbname,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Impossible d'ouvrir la DB : %v", err)
	}
	defer db.Close()

	// Vérifie la connexion
	if err := db.Ping(); err != nil {
		log.Fatalf("Échec du ping DB : %v", err)
	}

	e := echo.New()

	// Crée le serveur avec la connexion DB
	api := NewAPIServer(db)
	server.RegisterHandlers(e, api)

	// Démarre le serveur HTTP
	log.Println("Démarrage du serveur sur :8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Erreur lors du démarrage du serveur : %v", err)
	}
}
