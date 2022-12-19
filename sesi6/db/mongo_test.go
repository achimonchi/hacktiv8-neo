package db

import (
	"testing"
	"time"

	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/require"
	"gopkg.in/go-playground/mongostore.v4"
	mgo "gopkg.in/mgo.v2"
)

func TestConnect_Mongodb(t *testing.T) {
	client, err := NewMongoClient()
	require.Nil(t, err)
	require.NotNil(t, client)

}

func TestConnect_MongoSession(t *testing.T) {
	// Coonect to MongoDB
	dbSess, err := mgo.DialWithTimeout("mongodb://root:root@localhost:27017", 3*time.Second)
	if err != nil {
		panic(err)
	}
	defer dbSess.Close()

	store := mongostore.New(dbSess, "sessions", &sessions.Options{MaxAge: 3600}, true,
		[]byte("secret-key"))

	require.NotNil(t, store)
}
