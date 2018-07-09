package database

import "gopkg.in/mgo.v2"

func GetDatabaseConnection() *mgo.Session {
	session, err := mgo.Dial("")
	if err != nil {
		panic(err)
	}
	return session
}
