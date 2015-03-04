package paginate

import (
	"testing"

	"gopkg.in/nowk/assert.v2"
)

func TestPaginate(t *testing.T) {
	db, teardown := Setup(t)
	defer teardown()
	seed(t, db)

	pg := &Page{
		No:    2,
		Limit: 2,
	}

	var u Users
	qry := db.C(u.Collection()).Find(nil)
	err := Paginate(qry, pg, &u)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 3, pg.TotalPages())
	assert.Equal(t, 5, pg.TotalRecords())
	assert.Equal(t, 2, len(u))
}
