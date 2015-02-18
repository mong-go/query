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

Next/Prev pages

    page.NextPage() // New Page incremented or nil (if at the end of pages)
    page.PrevPage() // New Page decremented or nil (if at the first page)


---

__NewPagination__

Returns a `Paginated` object which contains the results, Page info and URL object to easily create pagination within your templates

    q := db.C("users").Find(bson.M{})
    var u []User
    p, err := query.NewPagination(q, &u, query.ParsePage(req), func(p *query.Paginated) {
      p.URL = &url.URL{
        Path:     "/users",
        RawQuery: "keyword=superheros",
      }
    })

    // p.Results       => interface{} (*[]User)
    // p.Page          => *query.Page (what is returned through Paginate)
    // p.PageURL()     => /users?keyword=superheros&page=2&per_page=30
    // p.NextPageURL() => /users?keyword=superheros&page=3&per_page=30
    // p.PrevPageURL() => /users?keyword=superheros&page=1&per_page=30

*`NextPageURL` and `PrevPageURL` will return `""` if there is no next/prev page.*


---

__ParsePage__

Small utility to return a `*Page` by parsing your URL querystring for `page` and `per_page` parameters.

    page := query.ParsePage(req)

    // page.No    => req.URL.Query().Get("page")
    // page.Limit => req.URL.Query().Get("per_page")


## License

MIT
