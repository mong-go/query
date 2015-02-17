package query

import (
	"gopkg.in/mgo.v2"
	"math"
	"reflect"
)

const (
	PerPageLimit = 30
)

type Page struct {
	// No is the current page
	No int

	// Limit is the per page limit
	Limit int

	count        int
	totalRecords int
}

func check(p *Page) *Page {
	if p.No == 0 {
		p.No = 1
	}

	if p.Limit == 0 {
		p.Limit = PerPageLimit
	}

	return p
}

func NewPage(n, lmt int) *Page {
	return check(&Page{
		No:    n,
		Limit: lmt,
	})
}

// Paginate maps an executed query to d and calculates pagination data returning
// it as Page
func Paginate(qry *mgo.Query, d interface{}, page *Page) (*Page, error) {
	p := *check(page)

	n, err := qry.Count()
	if err != nil {
		return nil, err
	}
	p.totalRecords = n

	if err := qry.Skip(p.Skip()).Limit(p.Limit).All(d); err != nil {
		return nil, err
	}
	v := reflect.ValueOf(d).Elem()
	p.count = v.Len()
	return &p, nil
}

func (p Page) Count() int {
	return p.count
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

func (p Page) Skip() int {
	if p.No == 1 {
		return 0
	}

	return (p.No - 1) * p.Limit
}

// NextPage returns a new Page with No incremented by 1 or returns nil if query
// as reached the last page
func (p Page) NextPage() *Page {
	if p.No*p.Limit >= p.totalRecords {
		return nil
	}

	return NewPage(p.No+1, p.Limit)
}

// PrevPage returns a new Page with No decremented by 1 or returns nil if at
// page no 1
func (p Page) PrevPage() *Page {
	if p.No < 2 {
		return nil
	}

	return NewPage(p.No-1, p.Limit)
}