package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"learn_verse/internal/db"
	"learn_verse/internal/router"
	"log"
	"time"
)

func main() {
	// 1. Config de la connexion (hardcodée ou à extraire en config/env)
	user := "postgres"
	password := "pwd"
	host := "localhost"
	port := "5432"
	dbname := "learn_verse"
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbname,
	)

	// 2. Connexion DB
	database, err := db.Connect(dsn)
	if err != nil {
		log.Fatalf("échec connexion DB : %v", err)
	}
	defer database.Close()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Setup(r, database)

	// 4. Lancement du serveur
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("échec serveur : %v", err)
	}
}
