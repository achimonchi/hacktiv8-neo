package db

import (
	"fmt"
	"testing"
)

func TestConnectPGStore(t *testing.T) {
	db, err := NewPostgres()
	if err != nil {
		t.Fatalf("error :%v", err.Error())

	}
	store, err := NewPostgresStore(db)
	if err != nil {
		t.Fatalf("error :%v", err.Error())
	}

	fmt.Printf("%+v\n", store)
}
