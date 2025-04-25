package db

import (
	"fmt"
	"testing"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pwd"
	dbname   = "learn_verse"
)

func TestConnect_Success(t *testing.T) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	database, err := Connect(dsn)
	if err != nil {
		t.Fatalf("La connexion devrait fonctionner, mais retourne une erreur : %v", err)
	}
	defer database.Close()
}
