package query

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mong-go/model.v1"
)

func Count(d model.ModelReader, b interface{}, db *mgo.Database) (int, error) {
	q, err := BuildQuery(d, b, db)
	if err != nil {
		return 0, err
	}

	return q.Count()
}
