package main

import (
	"database/sql"
	"encoding/json"
	"github.com/brianvoe/gofakeit/v6"
	"math/rand"
	"time"
)

func InsertFakeData(db *sql.DB) error {
	gofakeit.Seed(time.Now().UnixNano())

	var collectionIDs []string
	for i := 0; i < 3; i++ {
		name := gofakeit.Company()
		description := gofakeit.Sentence(10)
		var id string
		err := db.QueryRow(`
			INSERT INTO resource_collections (name, description)
			VALUES ($1, $2) RETURNING id
		`, name, description).Scan(&id)
		if err != nil {
			return err
		}
		collectionIDs = append(collectionIDs, id)
	}

	for i := 0; i < 10; i++ {
		title := gofakeit.Sentence(4)
		contentJSON, _ := json.Marshal(map[string]interface{}{"body": gofakeit.Paragraph(1, 2, 10, " ")})
		metadataJSON, _ := json.Marshal(map[string]interface{}{"author": gofakeit.Name(), "license": gofakeit.Word()})

		resourceType := resourceTypes[rand.Intn(len(resourceTypes))]
		collectionID := collectionIDs[rand.Intn(len(collectionIDs))]

		_, err := db.Exec(`
			INSERT INTO resources (collection_id, type, title, content, metadata)
			VALUES ($1, $2, $3, $4, $5)
		`, collectionID, resourceType, title, contentJSON, metadataJSON)

		if err != nil {
			return err
		}
	}

	return nil
}
