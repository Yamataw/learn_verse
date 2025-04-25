package main

import (
	"fmt"
	"learn_verse/internal/db"
	"testing"
)

func TestInsertFakeData(t *testing.T) {
	dsn := fmt.Sprintf("host=localhost port=5432 user=postgres password=pwd dbname=learn_verse sslmode=disable")

	conn, err := db.Connect(dsn)
	if err != nil {
		t.Fatalf("Connexion à la base échouée : %v", err)
	}
	defer conn.Close()

	err = InsertFakeData(conn)
	if err != nil {
		t.Errorf("Erreur lors de l'insertion des données factices : %v", err)
	}
}
