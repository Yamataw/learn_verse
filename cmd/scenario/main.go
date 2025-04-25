package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"learn_verse/client"
)

func main() {
	// Base URL de votre API OpenAPI 3.0.3
	const serverURL = "http://localhost:8080/api"

	ctx := context.Background()

	// 0) Instanciation du client typé avec réponses
	cli, err := client.NewClientWithResponses(serverURL)
	if err != nil {
		log.Fatalf("Erreur NewClientWithResponses: %v", err)
	}

	// --------------------------------------------------
	// 1) CREATE collection
	// --------------------------------------------------
	newColl := client.PostCollectionsJSONRequestBody{
		Name:        "MaCollectionGo",
		Description: ptrString("Créée via Go et oapi-codegen"),
	}
	createResp, err := cli.PostCollectionsWithResponse(ctx, newColl)
	if err != nil {
		log.Fatalf("PostCollectionsWithResponse failed: %v", err)
	}
	if createResp.StatusCode() != 201 {
		log.Fatalf("Échec création (status %d)", createResp.StatusCode())
	}
	created := createResp.JSON201
	fmt.Printf("✅ Créée : ID=%s, Name=%s, CreatedAt=%s\n",
		created.Id, created.Name, formatTimePtr(created.CreatedAt))

	// --------------------------------------------------
	// 2) LIST all collections
	// --------------------------------------------------
	listResp, err := cli.GetCollectionsWithResponse(ctx)
	if err != nil {
		log.Fatalf("GetCollectionsWithResponse failed: %v", err)
	}
	if listResp.StatusCode() != 200 {
		log.Fatalf("Échec listing (status %d)", listResp.StatusCode())
	}
	fmt.Println("📋 Liste des collections :")
	for _, c := range *listResp.JSON200 {
		fmt.Printf(" - %s : %s (updated at %s)\n",
			c.Id, c.Name, formatTimePtr(c.UpdatedAt))
	}

	// --------------------------------------------------
	// 3) GET collection by ID
	// --------------------------------------------------
	getResp, err := cli.GetCollectionsIdWithResponse(ctx, created.Id)
	if err != nil {
		log.Fatalf("GetCollectionsIdWithResponse failed: %v", err)
	}
	if getResp.StatusCode() != 200 {
		log.Fatalf("Échec GET by ID (status %d)", getResp.StatusCode())
	}
	fetched := getResp.JSON200
	fmt.Printf("🔍 Récupérée : ID=%s, Name=%s, Description=%s\n",
		fetched.Id, fetched.Name, ptrVal(fetched.Description))

	// --------------------------------------------------
	// 4) UPDATE collection (PUT)
	// --------------------------------------------------
	updateBody := client.PutCollectionsIdJSONRequestBody{
		Name:        "MaCollectionGo_Modifiée",
		Description: ptrString("Description mise à jour"),
	}
	putResp, err := cli.PutCollectionsIdWithResponse(ctx, created.Id, updateBody)
	if err != nil {
		log.Fatalf("PutCollectionsIdWithResponse failed: %v", err)
	}
	if putResp.StatusCode() != 200 {
		log.Fatalf("Échec update (status %d)", putResp.StatusCode())
	}
	updated := putResp.JSON200
	fmt.Printf("✏️  Mise à jour : Name=%s, UpdatedAt=%s\n",
		updated.Name, formatTimePtr(updated.UpdatedAt))

	// --------------------------------------------------
	// 5) DELETE collection
	// --------------------------------------------------
	delResp, err := cli.DeleteCollectionsIdWithResponse(ctx, created.Id)
	if err != nil {
		log.Fatalf("DeleteCollectionsIdWithResponse failed: %v", err)
	}
	if delResp.StatusCode() != 204 {
		log.Fatalf("Échec suppression (status %d)", delResp.StatusCode())
	}
	fmt.Println("🗑️  Supprimée avec succès.")
}

// ptrString retourne un pointeur vers s
func ptrString(s string) *string {
	return &s
}

// ptrVal lit la valeur de *string, ou "<nil>"
func ptrVal(s *string) string {
	if s == nil {
		return "<nil>"
	}
	return *s
}

// formatTimePtr gère les *time.Time nullable
func formatTimePtr(t time.Time) string {
	return t.Format(time.RFC3339)
}
