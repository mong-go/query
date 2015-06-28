package query

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mong-go/model.v1"
)

func All(d model.ModelReader, b interface{}, db *mgo.Database,
	opts ...QueryFunc) error {

	qry, err := BuildQuery(d, b, db, opts...)
	if err != nil {
		return err
	}

	return qry.All(d)
}
