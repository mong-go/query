package query

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mong-go/model.v0"
)

// One is a wrapper around mgo's One to supress "not found" error and return a
// boolean indicating "found" vs "not found".
func One(d model.ModelReader, b interface{}, db *mgo.Database,
	opts ...QueryFunc) (bool, error) {

	qry, err := BuildQuery(d, b, db, opts...)
	if err != nil {
		return false, err
	}

	if err := qry.One(d); err != nil {
		if err.Error() == "not found" {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
