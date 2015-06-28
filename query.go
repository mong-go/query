package query

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mong-go/model.v1"
)

type QueryFunc func(*mgo.Query) (*mgo.Query, error)

// BuildQuery compiles a query on the ModelReader collection
// Query for collection -> with filter -> in database [ -> with query config ]
func BuildQuery(d model.ModelReader, b interface{}, db *mgo.Database,
	opts ...QueryFunc) (*mgo.Query, error) {

	qry := db.C(d.CollectionName()).Find(b)
	for _, v := range opts {
		var err error
		if qry, err = v(qry); err != nil {
			return nil, err
		}
	}

	return qry, nil
}
