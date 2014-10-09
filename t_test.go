package query

import "testing"
import "gopkg.in/mgo.v2"

type User struct {
	Name string `name`
}

const databasename = "query_test"

func check(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("fatal: %s\n", err)
	}
}

func mongo(t *testing.T, fn func(*mgo.Database)) {
	session, err := mgo.Dial("127.0.0.1:27017")
	check(t, err)
	defer session.Close()

	db := session.DB(databasename)
	usersC := db.C("users")
	err = usersC.Insert(
		&User{"Batman"},
		&User{"Cobra Commander"})
	check(t, err)

	fn(db)

	err = usersC.Remove(nil)
	check(t, err)
}
