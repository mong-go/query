package query

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/nowk/assert.v2"
	"net/http"
	"net/url"
	"testing"
)

func seed(t *testing.T, db *mgo.Database) {
	i := 0
	for ; i < 5; i++ {
		err := UserFactory(db, func(u *User) {
			u.Name = fmt.Sprintf("User-%d", i+1)
		})
		if err != nil {
			t.Fatalf("factory error: %s", err)
		}
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
	} {
		var u []User
		qry := db.C("users").Find(bson.M{})
		page, err := Paginate(qry, &u, &Page{
			No:    v.No,
			Limit: v.Limit,
		})

		assert.Nil(t, err)
		assert.Equal(t, v.Current, page.No)
		assert.Equal(t, v.Limit, page.Limit)
		assert.Equal(t, v.Count, page.Count())
		assert.Equal(t, v.TotalRecords, page.TotalRecords())
		assert.Equal(t, v.TotalPages, page.TotalPages())
		assert.Equal(t, v.Skip, page.Skip())
		assert.Equal(t, page.Count(), len(u))
	}
}

func TestNextPage(t *testing.T) {
	{
		p := Page{
			No:           7,
			Limit:        4,
			totalRecords: 25,
		}
		assert.Nil(t, p.NextPage())
	}

	{
		p := Page{
			No:           6,
			Limit:        5,
			totalRecords: 30,
		}
		assert.Nil(t, p.NextPage())
	}

	{
		p := Page{
			No:           6,
			Limit:        5,
			totalRecords: 31,
		}
		assert.Equal(t, &Page{
			No:    7,
			Limit: 5,
		}, p.NextPage())
	}
}

func TestPrevPage(t *testing.T) {
	{
		p := Page{
			No: 1,
		}
		assert.Nil(t, p.PrevPage())
	}

	{
		p := Page{
			No: 0,
		}
		assert.Nil(t, p.PrevPage())
	}

	{
		p := Page{
			No:    2,
			Limit: 6,
		}
		assert.Equal(t, &Page{
			No:    1,
			Limit: 6,
		}, p.PrevPage())
	}
}

func TestParsePageReturnsPageFromRequestQueries(t *testing.T) {
	{
		req := &http.Request{
			URL: &url.URL{
				RawQuery: "page=2&per_page=12",
			},
		}

		page := ParsePage(req)
		assert.Equal(t, &Page{
			No:    2,
			Limit: 12,
		}, page)
	}

	{
		req := &http.Request{
			URL: &url.URL{},
		}

		page := ParsePage(req)
		assert.Equal(t, &Page{
			No:    1,
			Limit: 30,
		}, page)
	}
}

func TestNewPaginationReturnsPaginated(t *testing.T) {
	db, teardown := Setup(t)
	defer teardown()
	seed(t, db)

	var d []User
	qry := db.C("users").Find(bson.M{})
	p, err := NewPagination(qry, &d, &Page{
		No:    2,
		Limit: 3,
	}, func(p *Paginated) {
		p.URL = &url.URL{
			Path:     "/users/superhero",
			RawQuery: "foo=bar",
		}
	})
	assert.Nil(t, err)
	assert.Equal(t, 2, p.No)
	assert.Equal(t, 3, p.Limit)
	assert.Equal(t, 2, p.Count())
	assert.Equal(t, 5, p.TotalRecords())
	assert.Equal(t, 2, p.TotalPages())
	assert.Equal(t, p.Count(), len(*p.Results.(*[]User)))
	assert.Equal(t, "/users/superhero?foo=bar", p.URL.String())
	assert.Equal(t, "/users/superhero?foo=bar&page=2&per_page=3", p.PageURL())
	assert.Equal(t, "/users/superhero?foo=bar&page=1&per_page=3", p.PrevPageURL())
	assert.Equal(t, "", p.NextPageURL())
}
