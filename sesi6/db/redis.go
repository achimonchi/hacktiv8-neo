package db

import (
	"gopkg.in/boj/redistore.v1"
)

func NewRedisStore() (*redistore.RediStore, error) {
	authKey := []byte("authkey")
	encryptionKey := []byte("qMT53qxDCPmRhivTW7nLWVn3oLDdkiRT")

	store, err := redistore.NewRediStore(10, "tcp", "localhost:6379", "", authKey, encryptionKey)
	if err != nil {
		return nil, err
	}
	return store, nil
}
