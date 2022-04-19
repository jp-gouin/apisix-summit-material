# go-api

A mock REST API for AAA poc

## Create book

```bash
curl -X POST -H "Content-Type: application/json" \
    -d '{ "title" : "book1", "Author" : { "name" : "Bob", "lastName": "Bravo" }, "created" : "2009-11-10"}'\
    http://localhost:8080/api/book
```

## Get books

```bash
curl http://localhost:8080/api/books
```

## Create Author

```bash
curl -X POST -H "Content-Type: application/json" \
    -d '{ "name" : "Bob", "lastName": "Bravo" }'\
    http://localhost:8080/api/author
```

## Get Author

```bash
curl http://localhost:8080/api/authors
```
