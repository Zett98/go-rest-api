package sqlstore

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
)

// TestDB ... testing store
func TestDB(t *testing.T, databaseURL string) (*sql.DB, func(...string)) {
	t.Helper()
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		t.Fatal("error")
	}
	if err := db.Ping(); err != nil {
		t.Fatal("error")
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			_, err := db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
			if err != nil {
				t.Fatal(err)
			}
			err = db.Close()
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}
