package database

import (
	"testing"
)

const test_db_name = "test.db"

func TestNewDB(t *testing.T) {
	db := NewDB(test_db_name)

	if db == nil {
		t.Error(db)
	}

	db.Close()
}
