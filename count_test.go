package query

import (
	"testing"

	"gopkg.in/mgo.v2/bson"
	"gopkg.in/nowk/assert.v2"
)

func TestCount(t *testing.T) {
	db, teardown := Setup(t)
	defer teardown()

	for _, v := range []string{
		"Superman",
		"Batman",
		"Cobra Commander",
		"Captain Crunch",
	} {
		err := UserFactory(db, func(u *User) {
			u.Name = v
		})
		if err != nil {
			t.Fatal(err)
		}
	}

	var d Users
	n, err := Count(&d, bson.M{"name": bson.M{"$regex": "man"}}, db)
	assert.Nil(t, err)
	assert.Equal(t, 3, n)
}
