package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"
	"time"

	"learn_verse/client"
)

func main() {
	// Base URL de votre API en HTTPS
	const serverURL = "https://localhost:8443/api"

	// --- 1) Charger et configurer la CA pour TLS ---
	caCert, err := os.ReadFile("server.crt")
	if err != nil {
		log.Fatalf("Impossible de lire le certificat CA : %v", err)
	}
	caPool := x509.NewCertPool()
	if !caPool.AppendCertsFromPEM(caCert) {
		log.Fatal("Échec de l'ajout du certificat CA au pool")
	}

	// --- 2) Configuration TLS pour accepter les certificats sans SAN (fallback sur CN) ---
	tlsConfig := &tls.Config{
		// Utilise notre pool
		RootCAs: caPool,
		// Désactive la vérification standard (qui exige un SAN)
		InsecureSkipVerify: true,
		// Vérification manuelle : signature + chaîne, sans vérification de nom d'hôte
		VerifyPeerCertificate: func(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
			// Parse les certificats reçus
			certs := make([]*x509.Certificate, len(rawCerts))
			for i, asn1Data := range rawCerts {
				cert, err := x509.ParseCertificate(asn1Data)
				if err != nil {
					return err
				}
				certs[i] = cert
			}
			// Vérifie la chaîne contre notre pool de CA
			opts := x509.VerifyOptions{Roots: caPool}
			_, err := certs[0].Verify(opts)
			return err
		},
	}

	transport := &http.Transport{TLSClientConfig: tlsConfig}
	httpClient := &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second,
	}

	ctx := context.Background()

	// --- 3) Instanciation du client typé avec HTTP client TLS ---
	cli, err := client.NewClientWithResponses(serverURL, client.WithHTTPClient(httpClient))
	if err != nil {
		log.Fatalf("Erreur NewClientWithResponses: %v", err)
	}

	// --------------------------------------------------
	// 4) CREATE collection
	// --------------------------------------------------
	newColl := client.PostCollectionsJSONRequestBody{
		Name:        "MaCollectionGo",
		Description: ptrString("Créée via Go et oapi-codegen (TLS)"),
	}
	createResp, err := cli.PostCollectionsWithResponse(ctx, newColl)
	if err != nil {
		log.Fatalf("PostCollectionsWithResponse failed: %v", err)
	}
	if createResp.StatusCode() != http.StatusCreated {
		log.Fatalf("Échec création (status %d)", createResp.StatusCode())
	}
	created := createResp.JSON201
	log.Printf("✅ Créée : ID=%s, Name=%s, CreatedAt=%s", created.Id, created.Name, formatTimePtr(created.CreatedAt))

	// --------------------------------------------------
	// 5) LIST all collections
	// --------------------------------------------------
	listResp, err := cli.GetCollectionsWithResponse(ctx)
	if err != nil {
		log.Fatalf("GetCollectionsWithResponse failed: %v", err)
	}
	if listResp.StatusCode() != http.StatusOK {
		log.Fatalf("Échec listing (status %d)", listResp.StatusCode())
	}
	log.Println("📋 Liste des collections :")
	for _, c := range *listResp.JSON200 {
		log.Printf(" - %s : %s (updated at %s)", c.Id, c.Name, formatTimePtr(c.UpdatedAt))
	}

	// --------------------------------------------------
	// 6) GET collection by ID
	// --------------------------------------------------
	getResp, err := cli.GetCollectionsIdWithResponse(ctx, created.Id)
	if err != nil {
		log.Fatalf("GetCollectionsIdWithResponse failed: %v", err)
	}
	if getResp.StatusCode() != http.StatusOK {
		log.Fatalf("Échec GET by ID (status %d)", getResp.StatusCode())
	}
	fetched := getResp.JSON200
	log.Printf("🔍 Récupérée : ID=%s, Name=%s, Description=%s", fetched.Id, fetched.Name, ptrVal(fetched.Description))

	// --------------------------------------------------
	// 7) UPDATE collection (PUT)
	// --------------------------------------------------
	updateBody := client.PutCollectionsIdJSONRequestBody{
		Name:        "MaCollectionGo_Modifiée",
		Description: ptrString("Description mise à jour (TLS)"),
	}
	putResp, err := cli.PutCollectionsIdWithResponse(ctx, created.Id, updateBody)
	if err != nil {
		log.Fatalf("PutCollectionsIdWithResponse failed: %v", err)
	}
	if putResp.StatusCode() != http.StatusOK {
		log.Fatalf("Échec update (status %d)", putResp.StatusCode())
	}
	updated := putResp.JSON200
	log.Printf("✏️  Mise à jour : Name=%s, UpdatedAt=%s", updated.Name, formatTimePtr(updated.UpdatedAt))

	// --------------------------------------------------
	// 8) DELETE collection
	// --------------------------------------------------
	delResp, err := cli.DeleteCollectionsIdWithResponse(ctx, created.Id)
	if err != nil {
		log.Fatalf("DeleteCollectionsIdWithResponse failed: %v", err)
	}
	if delResp.StatusCode() != http.StatusNoContent {
		log.Fatalf("Échec suppression (status %d)", delResp.StatusCode())
	}
	log.Println("🗑️  Supprimée avec succès.")
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

// formatTimePtr formate *time.Time
func formatTimePtr(t time.Time) string {
	return t.Format(time.RFC3339)
}
