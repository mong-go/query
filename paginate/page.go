package paginate

import (
	"math"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2"
)

type Page struct {
	No    int
	Limit int

	totalRecords int
}

func NewPage(no, lmt int) *Page {
	return checkPage(&Page{
		No:    no,
		Limit: lmt,
	})
}

func (p Page) TotalRecords() int {
	return p.totalRecords
}

func (p Page) TotalPages() int {
	if p.totalRecords == 0 {
		return 0
	}
	if p.totalRecords <= p.Limit {
		return 1
	}
	return int(math.Ceil(float64(p.totalRecords) / float64(p.Limit)))
}

// HasPages returns a bool based on if there is more than 1 page.
func (p Page) HasPages() bool {
	return p.TotalPages() > 1
}

func (p Page) Skip() int {
	if p.No == 1 {
		return 0
	}
	return (p.No - 1) * p.Limit
}

func (p Page) NextPage() *Page {
	if p.No*p.Limit >= p.totalRecords {
		return nil
	}
	return NewPage(p.No+1, p.Limit)
}

func (p Page) PrevPage() *Page {
	if p.No < 2 {
		return nil
	}
	return NewPage(p.No-1, p.Limit)
}

// Query is a short for Query bound to this Page
func (p *Page) Query(qry *mgo.Query) (*mgo.Query, error) {
	return Query(qry, p)
}

// PerPageLimit is the default per/page limit
const PerPageLimit = 30

// checkPage checks that a page is at least 1 and limit is > 0
func checkPage(p *Page) *Page {
	if p.No <= 0 {
		p.No = 1
	}
	if p.Limit <= 0 {
		p.Limit = PerPageLimit
	}
	return p
}

// ParsePage is a helper to return a page by parsing url query values
// eg. http://example.com/posts?page=2&per_page=100
func ParsePage(req *http.Request) *Page {
	q := req.URL.Query()
	n, _ := strconv.ParseInt(q.Get("page"), 10, 64)
	l, _ := strconv.ParseInt(q.Get("per_page"), 10, 64)

	return checkPage(&Page{
		No:    int(n),
		Limit: int(l),
	})
}
