package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

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
		log.Fatalf("❌ Erreur à l'ouverture de la connexion : %v", err)
	}

	db.SetMaxOpenConns(10)                 // nb max de connexions ouvertes
	db.SetMaxIdleConns(5)                  // nb max de connexions inactives
	db.SetConnMaxIdleTime(5 * time.Minute) // durée max d’inactivité
	db.SetConnMaxLifetime(1 * time.Hour)   // durée max de vie d’une connexion

	// 5. Vérification de la connexion via un ping
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("❌ Impossible de joindre PostgreSQL : %v", err)
	}
	fmt.Println("✅ Connecté à PostgreSQL !")

	// 6. Exemple de requête simple
	var now time.Time
	err = db.QueryRowContext(ctx, "SELECT NOW()").Scan(&now)
	
	if err != nil {
		log.Fatalf("❌ Erreur lors de la requête : %v", err)
	}
	fmt.Printf("Heure actuelle selon PostgreSQL : %s\n", now.Format(time.RFC3339))
}
