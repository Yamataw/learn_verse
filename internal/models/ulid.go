// models/ulid.go
package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/oklog/ulid/v2"
)

// ULID wrappe ulid.ULID pour supporter SQL et JSON.
type ULID ulid.ULID

// Scan implémente sql.Scanner pour lire du TEXT (26 bytes) ou du BINARY (16 bytes).
func (u *ULID) Scan(value interface{}) error {
	var data []byte
	switch v := value.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	default:
		return fmt.Errorf("models.ULID.Scan: type %T non supporté", value)
	}

	switch len(data) {
	case ulid.EncodedSize: // 26 caractères Base32
		if err := (*ulid.ULID)(u).UnmarshalText(data); err != nil {
			return fmt.Errorf("models.ULID.UnmarshalText: %w", err)
		}
		return nil
	default:
		return fmt.Errorf("models.ULID.Scan: taille inattendue %d", len(data))
	}
}

// Value implémente driver.Valuer pour convertir ULID en chaîne Base32 à l’insertion.
func (u ULID) Value() (driver.Value, error) {
	return (*ulid.ULID)(&u).String(), nil
}

// MarshalText permet d’exporter ULID en Base32 (pour JSON et autres encodages texte).
func (u ULID) MarshalText() ([]byte, error) {
	return []byte((*ulid.ULID)(&u).String()), nil
}

// UnmarshalText parse une chaîne Base32 en ULID.
func (u *ULID) UnmarshalText(data []byte) error {
	parsed, err := ulid.Parse(string(data))
	if err != nil {
		return err
	}
	*u = ULID(parsed)
	return nil
}

// MarshalJSON appelle automatiquement MarshalText et l’enveloppe en JSON string.
func (u ULID) MarshalJSON() ([]byte, error) {
	text, err := u.MarshalText()
	if err != nil {
		return nil, err
	}
	return json.Marshal(string(text))
}

// UnmarshalJSON extrait le string du JSON puis appelle UnmarshalText.
func (u *ULID) UnmarshalJSON(data []byte) error {
	// on attend un JSON string
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	return u.UnmarshalText([]byte(s))
}
