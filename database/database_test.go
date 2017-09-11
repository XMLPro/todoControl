package database

import "testing"

func TestDB_GetUserPass(t *testing.T) {
	db := NewDB("../test.db")
	pass := db.GetUserPass("test_user2")
	if pass != ""{

	}
}