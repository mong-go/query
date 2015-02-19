package query

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mong-go/model.v0"
)

// One is a wrapper around mgo's One to supress "not found" error and return a
// boolean indicating "found" vs "not found".
func One(db *mgo.Database, d model.ModelReader, b bson.M, opts ...QueryFunc) (bool, error) {
	qry := db.C(d.Collection()).Find(b)
	for _, v := range opts {
		v(qry)
	}

	err := qry.One(d)
	if err == nil {
		return true, nil
	}
	if err.Error() == "not found" {
		return false, nil
	}

	return false, err
}
