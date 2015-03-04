package paginate

import (
	"fmt"
	"testing"

	"gopkg.in/mgo.v2"
	"gopkg.in/mong-go/mongod.v1"
	"gopkg.in/mong-go/query.v2"
	"gopkg.in/nowk/assert.v2"
)

func Setup(t *testing.T) (*mgo.Database, func()) {
	m := mongod.New("query_test")
	db, err := m.Start()
	if err != nil {
		t.Fatal(err)
	}

	return db, func() {
		m.Stop(func(db *mgo.Database) {
			_, err := db.C("users").RemoveAll(nil)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

type User struct {
	Name string
}

func (User) Collection() string {
	return "users"
}

type Users []User

func (Users) Collection() string {
	return "users"
}

func seed(t *testing.T, db *mgo.Database) {
	i := 0
	for ; i < 5; i++ {
		db.C("users").Insert(User{
			Name: fmt.Sprintf("User %d", i),
		})
	}
}

func TestQuery(t *testing.T) {
	db, teardown := Setup(t)
	defer teardown()
	seed(t, db)

	for _, v := range []struct {
		No    int
		Limit int

		Count        int
		TotalRecords int
		TotalPages   int
		Skip         int
	}{
		{0, 2 /**/, 2, 5, 3, 0},
		{1, 2 /**/, 2, 5, 3, 0},
		{2, 2 /**/, 2, 5, 3, 2},
		{3, 2 /**/, 1, 5, 3, 4},
		{1, 5 /**/, 5, 5, 1, 0},
	} {
		pg := NewPage(v.No, v.Limit)

		var u Users
		err := query.All(&u, nil, db, func(qry *mgo.Query) (*mgo.Query, error) {
			return Query(qry, pg)
		})
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, v.TotalRecords, pg.TotalRecords())
		assert.Equal(t, v.TotalPages, pg.TotalPages())
		assert.Equal(t, v.Skip, pg.Skip())
		assert.Equal(t, v.Count, len(u))
	}
}
