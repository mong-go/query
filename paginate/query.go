package paginate

import (
	"gopkg.in/mgo.v2"
)

func Query(qry *mgo.Query, page *Page) (*mgo.Query, error) {
	page = checkPage(page)

	n, err := qry.Count()
	if err != nil {
		return nil, err
	}

	page.totalRecords = n

	return qry.Skip(page.Skip()).Limit(page.Limit), nil
}
