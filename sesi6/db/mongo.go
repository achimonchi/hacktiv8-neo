package db

import (
	"log"

	"gopkg.in/mgo.v2"
)

func NewMongoSession() *mgo.Session {
	// uri := "mongodb+srv://hacktiv8:hacktiv8@hacktiv8.9kzvpvd.mongodb.net/?ssl=true"
	uri := "mongodb://localhost:27017"

	session, err := mgo.Dial(uri)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return session
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
