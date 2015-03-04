# query

[![Build Status](https://travis-ci.org/mong-go/query.svg?branch=master)](https://travis-ci.org/mong-go/query) [![GoDoc](https://godoc.org/gopkg.in/mong-go/query.v2?status.svg)](http://godoc.org/gopkg.in/mong-go/query.v2)

mgo query utilities

## Install

    go get gopkg.in/mong-go/query.v2

## Usage

One/All funcs require `mong-go/model`'s `ModelReader` interface.

    type User struct {
      Name string
    }

    func (User) Collection() String {
      return "users"
    }

__One__

`One` is a wrapper around `mgo`'s own `One` func which supresses the `not found` error and provides a `bool` indicating *found* or *not found*.

    var u User
    ok, err := query.One(&u, bson.M{"name": "Batman"}, db)
    if err != nil {
      // handle error
    }
    if !ok {
      // handle not found
    }

---

__All__

`All` is a general wrapper around `mgo`'s own `All` func.

    type Users []User

    func (Users) Collection() string {
      return "users"
    }

    var u Users
    err := query.All(&u, bson.M{}, db)
    if err != nil {
      // handle error
    }

---

#### /paginate

The `paginate` package provides utilities for pagination.

    import "gopkg.in/mong-go/query.v2/paginate"

    pg := &paginate.Page{
      No: 2,
      Limit: 30,
    }

    var u Users
    err := query.All(&u, bson.M{}, db, func(qry *mgo.Query) (*mgo.Query, error) {
      return paginate.Query(qry.Sort("+name"), pg)
    })

    // pg.TotalRecords() => totals returned for the query
    // pg.TotalPages()   => total pages

Next/Prev pages

    pg.NextPage() // New Page incremented or nil (if at the end of pages)
    pg.PrevPage() // New Page decremented or nil (if at the first page)

---

__ParsePage__

Small utility to return a `*Page` by parsing your URL querystring for `page` and `per_page` parameters.

    http://example.com/posts?page=3&per_page=50

Parsing results in

    pg := paginate.ParsePage(req.URL)

    // pg.No    => 3
    // pg.Limit => 50


## License

MIT
