package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	mgo "gopkg.in/mgo.v2"
	// "labix.org/v2/mgo"
)

func NewMongoSession() *mgo.Session {
	// uri := "mongodb+srv://hacktiv8:hacktiv8@hacktiv8.9kzvpvd.mongodb.net/?authSource=admin"
	uri := "mongodb://root:root@localhost:27017"
	// uri := "localhost:27017"

	// maxWait := time.Duration(5 * time.Second)
	// Connect to the MongoDB server
	session, err := mgo.Dial(uri)
	if err != nil {
		panic(err)
	}

	return session

	// return session
}

func NewMongoClient() (*mongo.Client, error) {
	uri := "mongodb+srv://hacktiv8:hacktiv8@hacktiv8.9kzvpvd.mongodb.net/?authSource=admin"
	// uri := "mongodb://root:root@localhost:27017"

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return client, nil
}

// func ConnectMongoStore(session *mgo.Session, db, collection string) *mongostore.MongoStore {
// 	dbCollection := session.DB(db).C(collection)
// 	maxAge := 20
// 	ensureTTL := true
// 	authKey := []byte("authkey")
// 	encryptionKey := []byte("qMT53qxDCPmRhivTW7nLWVn3oLDdkiRT")

// 	store := mongostore.NewMongoStore(
// 		dbCollection,
// 		maxAge,
// 		ensureTTL,
// 		authKey, encryptionKey,
// 	)

// 	return store
// }
