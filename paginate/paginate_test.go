package paginate

import (
	"fmt"
	"testing"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mong-go/mongod.v1"
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
			_, err := db.C("users").RemoveAll(bson.M{})
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func seed(t *testing.T, db *mgo.Database) {
	i := 0
	for ; i < 5; i++ {
		db.C("users").Insert(map[string]interface{}{
			"name": fmt.Sprintf("User %d", i),
		})
	}
}

func TestPaginate(t *testing.T) {
	db, teardown := Setup(t)
	defer teardown()
	seed(t, db)

	for _, v := range []struct {
		No    int
		Limit int

		Current      int
		Count        int
		TotalRecords int
		TotalPages   int
		Skip         int
	}{
		{0, 2 /**/, 1, 2, 5, 3, 0},
		{1, 2 /**/, 1, 2, 5, 3, 0},
		{2, 2 /**/, 2, 2, 5, 3, 2},
		{3, 2 /**/, 3, 1, 5, 3, 4},
		{1, 5 /**/, 1, 5, 5, 1, 0},
	} {
		page := NewPage(v.No, v.Limit)

		var u []map[string]interface{}
		err := Paginate(db.C("users").Find(bson.M{}), &u, page)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, v.Current, page.No)
		assert.Equal(t, v.Limit, page.Limit)
		assert.Equal(t, v.Count, page.Count())
		assert.Equal(t, v.TotalRecords, page.TotalRecords())
		assert.Equal(t, v.TotalPages, page.TotalPages())
		assert.Equal(t, v.Skip, page.Skip())
		assert.Equal(t, page.Count(), len(u))
	}
}
