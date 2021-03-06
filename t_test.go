package query

import (
	"testing"

	"gopkg.in/mgo.v2"
	"gopkg.in/mong-go/model.v0"
	"gopkg.in/mong-go/mongod.v1"
)

type User struct {
	Name string `name`
}

func (User) Collection() string {
	return "users"
}

var _ model.ModelReader = &User{}

const databasename = "query_test"

func cleandb(t *testing.T) func(*mgo.Database) {
	return func(db *mgo.Database) {
		if _, err := db.C("users").RemoveAll(nil); err != nil {
			t.Fatal(err)
		}
	}
}

func Setup(t *testing.T) (*mgo.Database, func()) {
	m := mongod.New(databasename)
	db, err := m.Start()
	if err != nil {
		t.Fatal(err)
	}

	return db, func() {
		m.Stop(cleandb(t))
	}
}

func UserFactory(db *mgo.Database, opts ...func(*User)) error {
	u := &User{
		Name: "Superman",
	}
	for _, v := range opts {
		v(u)
	}

	return db.C("users").Insert(u)
}
