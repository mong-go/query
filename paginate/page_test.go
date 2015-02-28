package paginate

import (
	"net/http"
	"net/url"
	"testing"

	"gopkg.in/nowk/assert.v2"
)

func TestPageMaths(t *testing.T) {
	for _, v := range []struct {
		No           int
		Limit        int
		Count        int
		TotalRecords int

		tTotalPages int
		tSkip       int
	}{
		{1, 2, 2, 5 /**/, 3, 0},
		{2, 2, 2, 5 /**/, 3, 2},
		{3, 2, 1, 5 /**/, 3, 4},
		{1, 8, 8, 8 /**/, 1, 0}, // query count matches page limit
		{1, 5, 0, 0 /**/, 0, 0}, // no results
	} {
		page := Page{
			No:    v.No,
			Limit: v.Limit,

			count:        v.Count,
			totalRecords: v.TotalRecords,
		}

		assert.Equal(t, v.tTotalPages, page.TotalPages())
		assert.Equal(t, v.tSkip, page.Skip())
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
