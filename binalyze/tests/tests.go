package tests

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var _noresult int
var NoResult interface{} = &_noresult

func TestOne(t *testing.T, query string, expected interface{}) {
	t.Helper()
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		t.Fatal(err)
	}

	defer rows.Close()
	count := 0
	expectedCount := 1
	if expected == NoResult {
		expectedCount = 0
	}
	for rows.Next() {
		count++
		var result interface{}
		err = rows.Scan(&result)
		if err != nil {
			t.Fatal(err)
		}
		if result != expected {
			t.Fatalf("expected %v(%[1]T), got %v(%[2]T)", expected, result)
		}
	}
	if count != expectedCount {
		t.Fatalf("expected %d row, got %d", expectedCount, count)
	}
}
