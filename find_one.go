package query

import "gopkg.in/mgo.v2"

// FindOne is a wrapper around mgo's One to supress "not found" error and return
// a boolean indicating "found" vs "not found".
// All other errors will be returned
func FindOne(q *mgo.Query, d interface{}) (bool, error) {
	err := q.One(d)
	if err == nil {
		return true, nil
	}

	if err.Error() == "not found" {
		return false, nil
	}

	return false, err
}
