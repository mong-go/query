package query

import (
	"gopkg.in/mgo.v2"
)

type QueryFunc func(*mgo.Query) *mgo.Query
