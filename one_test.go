package query

import "testing"
import "gopkg.in/mgo.v2/bson"
import "gopkg.in/nowk/assert.v2"

func TestOneFound(t *testing.T) {
	db, teardown := Setup(t)
	defer teardown()
	if err := UserFactory(db, func(u *User) {
		u.Name = "Batman"
	}); err != nil {
		t.Fatal(err)
	}

	var user User
	ok, err := One(db, &user, bson.M{"name": "Batman"})
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestOneNotFound(t *testing.T) {
	db, teardown := Setup(t)
	defer teardown()

	var user User
	ok, err := One(db, &user, bson.M{"name": "Robin"})
	assert.Nil(t, err)
	assert.False(t, ok)
}
