# query

[![Build Status](https://travis-ci.org/gomon/query.svg?branch=master)](https://travis-ci.org/gomon/query)
[![GoDoc](https://godoc.org/github.com/gomon/query?status.svg)](http://godoc.org/github.com/gomon/query)

mgo query utilities

## Examples

`FindOne` supresses the `not found` error and provides a `bool` indicating *found* or *not found*.

    qry := users.Find(bson.M{"name": "Batman"})
    found, err := query.FindOne(qry, &user)
    if err != nil {
      // handle error
    }

    if found {
      // handle found
      return
    }

    // handle not found

## License

MIT
