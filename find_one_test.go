package query

import "testing"
import "gopkg.in/mgo.v2/bson"
import "gopkg.in/nowk/assert.v2"

func TestFindOneFound(t *testing.T) {
	db, teardown := Setup(t)
	defer teardown()
	if err := UserFactory(db, func(u *User) {
		u.Name = "Batman"
	}); err != nil {
		t.Fatal(err)
	}

	var user User
	users := db.C("users")
	qry := users.Find(bson.M{"name": "Batman"})
	found, err := FindOne(qry, &user)
	assert.True(t, found)
	assert.Nil(t, err)
}

func TestFindOneNotFound(t *testing.T) {
	db, teardown := Setup(t)
	defer teardown()

	var user User
	users := db.C("users")
	qry := users.Find(bson.M{"name": "Robin"})
	found, err := FindOne(qry, &user)
	assert.False(t, found)
	assert.Nil(t, err)
}
