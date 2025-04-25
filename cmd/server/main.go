package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/unrolled/secure"
	"learn_verse/internal/db"
	"learn_verse/internal/router"
)

func main() {
	// === 1. Connexion à la DB ===
	user := "postgres"
	password := "pwd"
	host := "localhost"
	port := "5432"
	dbname := "learn_verse"
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbname,
	)
	database, err := db.Connect(dsn)
	if err != nil {
		log.Fatalf("échec connexion DB : %v", err)
	}
	defer database.Close()

	// === 2. Setup Gin + CORS ===
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

	// === 3. Middleware de sécurité pour forcer HTTPS ===
	secureMiddleware := secure.New(secure.Options{
		SSLRedirect:          true,
		SSLHost:              "localhost:8443", // host:port où tourne le HTTPS
		SSLProxyHeaders:      map[string]string{"X-Forwarded-Proto": "https"},
		STSSeconds:           31536000, // HSTS one year
		STSIncludeSubdomains: true,
		FrameDeny:            true,
		ContentTypeNosniff:   true,
		BrowserXssFilter:     true,
	})
	// On wrappe le router Gin dans la couche de redirection
	handler := secureMiddleware.Handler(r)

	// === 4. Lancement du serveur HTTP (redirection) ===
	go func() {
		httpAddr := ":8080"
		log.Printf("Démarrage du serveur HTTP (redirection vers HTTPS) sur %s", httpAddr)
		// toute requête HTTP sera redirigée vers HTTPS://localhost:8443/...
		if err := http.ListenAndServe(httpAddr, handler); err != nil {
			log.Fatalf("échec serveur HTTP : %v", err)
		}
	}()

	// === 5. Configuration TLS personnalisée ===
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
	}

	// === 6. Lancement du serveur HTTPS ===
	httpsSrv := &http.Server{
		Addr:         ":8443",
		Handler:      handler, // même middleware + router
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSConfig:    tlsConfig,
	}
	log.Printf("Démarrage du serveur HTTPS sur %s", httpsSrv.Addr)
	if err := httpsSrv.ListenAndServeTLS("server.crt", "server.key"); err != nil {
		log.Fatalf("échec serveur HTTPS : %v", err)
	}
}
