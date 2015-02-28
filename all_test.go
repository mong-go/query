package query

import (
	"testing"

	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mong-go/model.v0"
	"gopkg.in/nowk/assert.v2"
)

type Users []User

func (Users) Collection() string {
	return "users"
}

var _ model.ModelReader = Users{}

func TestAll(t *testing.T) {
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
	err := All(&d, bson.M{"name": bson.M{"$regex": "man"}}, db)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(d))
}
