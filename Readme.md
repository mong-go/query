# query

[![Build Status](https://travis-ci.org/mong-go/query.svg?branch=master)](https://travis-ci.org/mong-go/query)
[![GoDoc](https://godoc.org/gopkg.in/mong-go/query.v1?status.svg)](http://godoc.org/gopkg.in/mong-go/query.v1)

mgo query utilities

## Install

    go get gopkg.in/mong-go/query.v1

## Usage

__FindOne__

Supresses the `not found` error and provides a `bool` indicating *found* or *not found*.

    q := db.C("users").Find(bson.M{"name": "Batman"})
    ok, err := query.FindOne(q, &user)
    if err != nil {
      // handle error
    }

    if ok {
      // handle found
      return
    }

    // handle not found

---

__Paginate__

Provides functionality for pagination.

    q := db.C("users").Find(bson.M{})
    var u []User
    page, err := query.Paginate(q, &u, &query.Page{
      No:    1,
      Limit: 30,
    })

    // page.Count()        => totals returned for the given page
    // page.TotalRecords() => totals returned for the query
    // page.TotalPages()   => total pages

Query the next or previous pages

    p2, err := query.Paginate(q, &u, p1.NextPage())

    p1, err := query.Paginate(q, &u, p2.PrevPage())
    

## License

MIT
