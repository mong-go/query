package query

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mong-go/model.v0"
)

func All(d model.ModelReader, b bson.M, db *mgo.Database,
	opts ...QueryFunc) error {

	qry, err := BuildQuery(d, b, db, opts...)
	if err != nil {
		return err
	}

	return qry.All(d)
}
