package paginate

import (
	"reflect"

	"gopkg.in/mgo.v2"
)

// Paginate queries all on to d @ Page
func Paginate(qry *mgo.Query, d interface{}, page *Page) error {
	var err error
	qry, err = page.Query(qry)
	if err != nil {
		return err
	}
	if err := qry.All(d); err != nil {
		return err
	}

	page.count = reflect.ValueOf(d).Elem().Len()

	return nil
}
