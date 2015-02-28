package query

import (
	"testing"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mong-go/model.v0"
	"gopkg.in/mong-go/query.v2/paginate"
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

func TestAllWithPaginate(t *testing.T) {
	db, teardown := Setup(t)
	defer teardown()

	for _, v := range []string{
		"1",
		"2",
		"3",
		"4",
	} {
		err := UserFactory(db, func(u *User) {
			u.Name = v
		})
		if err != nil {
			t.Fatal(err)
		}
	}

	page := paginate.NewPage(1, 2)

	var d Users
	err := All(&d, bson.M{}, db, func(qry *mgo.Query) (*mgo.Query, error) {
		return page.Query(qry)
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 2, len(d))
}
