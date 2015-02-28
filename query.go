package query

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mong-go/model.v0"
)

type QueryFunc func(*mgo.Query) (*mgo.Query, error)

// BuildQuery compiles a query on the ModelReader collection
// Query for collection -> with filter -> in database [ -> with query config ]
func BuildQuery(d model.ModelReader, b bson.M, db *mgo.Database,
	opts ...QueryFunc) (*mgo.Query, error) {

	qry := db.C(d.Collection()).Find(b)
	for _, v := range opts {
		var err error
		if qry, err = v(qry); err != nil {
			return nil, err
		}
	}

	return qry, nil
}
