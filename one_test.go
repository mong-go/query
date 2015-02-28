package query

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/nowk/assert.v2"
	"testing"
)

func TestOneFound(t *testing.T) {
	db, teardown := Setup(t)
	defer teardown()
	if err := UserFactory(db, func(u *User) {
		u.Name = "Batman"
	}); err != nil {
		t.Fatal(err)
	}

	var user User
	ok, err := One(&user, bson.M{"name": "Batman"}, db)
	assert.Nil(t, err)
	assert.True(t, ok)
}

func TestOneNotFound(t *testing.T) {
	db, teardown := Setup(t)
	defer teardown()
	if err := UserFactory(db, func(u *User) {
		u.Name = "Batman"
	}); err != nil {
		t.Fatal(err)
	}

	var user User
	ok, err := One(&user, bson.M{"name": "Robin"}, db)
	assert.Nil(t, err)
	assert.False(t, ok)
}
