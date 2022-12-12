package db

import (
	"testing"
)

func TestConnect_Mongodb(t *testing.T) {
	session := NewMongoSession()
	if session == nil {
		t.Errorf("error")
	}
}
