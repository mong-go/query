package query

import "github.com/gomon/mongod"
import "testing"
import "gopkg.in/mgo.v2/bson"
import "gopkg.in/nowk/assert.v2"

func TestFindOneFound(t *testing.T) {
	m := mongod.New(databasename)
	db, err := m.Start()
	if err != nil {
		t.Fatal(err)
	}
	defer m.Stop(cleandb)

	var user User
	users := db.C("users")
	qry := users.Find(bson.M{"name": "Batman"})
	found, err := FindOne(qry, &user)
	assert.True(t, found)
	assert.Nil(t, err)
}

func TestFindOneNotFound(t *testing.T) {
	m := mongod.New(databasename)
	db, err := m.Start()
	if err != nil {
		t.Fatal(err)
	}
	defer m.Stop(cleandb)

	var user User
	users := db.C("users")
	qry := users.Find(bson.M{"name": "Robin"})
	found, err := FindOne(qry, &user)
	assert.False(t, found)
	assert.Nil(t, err)
}
