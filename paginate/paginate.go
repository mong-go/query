package paginate

import (
	"gopkg.in/mgo.v2"
)

// Paginate paginates on query
func Paginate(qry *mgo.Query, pg *Page, d interface{}) error {
	qry, err := Query(qry, pg)
	if err != nil {
		return err
	}

	err = qry.All(d)
	if err != nil {
		return err
	}

	return nil
}
