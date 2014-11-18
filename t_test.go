package query

import "gopkg.in/mgo.v2"

type User struct {
	Name string `name`
}

const databasename = "query_test"

var cleandb = func(db *mgo.Database) {
	u := db.C("users")
	u.Remove(nil)
}
