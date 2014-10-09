package query

import "testing"
import "gopkg.in/mgo.v2"
import "gopkg.in/mgo.v2/bson"
import "github.com/nowk/assert"

func TestFindOneFound(t *testing.T) {
	mongo(t, func(db *mgo.Database) {
		var user User
		users := db.C("users")
		qry := users.Find(bson.M{"name": "Batman"})
		found, err := FindOne(qry, &user)
		assert.True(t, found)
		assert.Nil(t, err)
	})
}

func TestFindOneNotFound(t *testing.T) {
	mongo(t, func(db *mgo.Database) {
		var user User
		users := db.C("users")
		qry := users.Find(bson.M{"name": "Robin"})
		found, err := FindOne(qry, &user)
		assert.False(t, found)
		assert.Nil(t, err)
	})
}
