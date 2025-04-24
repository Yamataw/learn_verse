package main

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	_ "github.com/lib/pq"
	"learn_verse/internal/db"
	"log"
	"math/rand"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pwd"
	dbname   = "learn_verse"
)

var resourceTypes = []string{"note", "flashcard", "quiz", "file"}

func main() {
	gofakeit.Seed(time.Now().UnixNano())

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// 🔄 Connexion via la fonction utilitaire
	database, err := db.Connect(dsn)
	if err != nil {
		log.Fatalf("échec connexion DB : %v", err)
	}
	defer database.Close()

	// Insertion resource_collections
	var collectionIDs []string
	for i := 0; i < 3; i++ {
		name := gofakeit.Company()
		description := gofakeit.Sentence(10)
		var id string
		err := database.QueryRow(`
			INSERT INTO resource_collections (name, description)
			VALUES ($1, $2) RETURNING id
		`, name, description).Scan(&id)
		if err != nil {
			log.Fatalf("Erreur insert collection: %v", err)
		}
		collectionIDs = append(collectionIDs, id)
	}

	// Insertion resources
	for i := 0; i < 10; i++ {
		title := gofakeit.Sentence(4)
		contentMap := map[string]interface{}{
			"body": gofakeit.Paragraph(1, 2, 10, " "),
		}
		contentJSON, _ := json.Marshal(contentMap)

		metadataMap := map[string]interface{}{
			"author":  gofakeit.Name(),
			"license": gofakeit.Word(),
		}
		metadataJSON, _ := json.Marshal(metadataMap)

		resourceType := resourceTypes[rand.Intn(len(resourceTypes))]
		collectionID := collectionIDs[rand.Intn(len(collectionIDs))]

		_, err := database.Exec(`
			INSERT INTO resources (collection_id, type, title, content, metadata)
			VALUES ($1, $2, $3, $4, $5)
		`, collectionID, resourceType, title, contentJSON, metadataJSON)
		if err != nil {
			log.Fatalf("Erreur insert resource: %v", err)
		}
	}

	fmt.Println("🎉 Données de test insérées avec succès !")
}
