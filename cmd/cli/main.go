package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"learn_verse/client"
)

func main() {
	serverURL := "http://localhost:8080/api"

	// Crée un client simple (sans helper pour les réponses typées)
	cli, err := client.NewClient(serverURL)
	if err != nil {
		log.Fatalf("Erreur lors de la création du client : %v", err)
	}

	// Contexte pour les appels HTTP
	ctx := context.Background()

	// Appelle GET /collections
	resp, err := cli.GetCollections(ctx)
	if err != nil {
		log.Fatalf("Erreur lors de l'appel à GetCollections : %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Statut inattendu : %s", resp.Status)
	}

	// Décode la réponse JSON
	var cols []client.ResourceCollection
	if err := json.NewDecoder(resp.Body).Decode(&cols); err != nil {
		log.Fatalf("Impossible de décoder le JSON : %v", err)
	}

	// Affiche chaque collection
	for _, col := range cols {
		fmt.Printf("ID: %s\tName: %s\n", col.Id, col.Name)
	}
}
