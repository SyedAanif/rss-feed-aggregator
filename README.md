# rss-feed-aggregator
This repository contains a service written in GO for RSS(RDF Site Summary or Really Simple Syndication) feed aggregation.

# Concepts Covered:
- GO Lang
- Postgres
- Web-Server using CHI
- Thunder Client extension of VSCode

# Commands and packages used:
- `go mod init`
- `go mod tidy`
- `go mod vendor`
- `go get github.com/go-chi/chi`
- `go get github.com/go-chi/cors`
- `go get github.com/google/uuid`
- `go get github.com/joho/godotenv`
- `go get github.com/lib/pq`

```
brew install postgresql
brew services start postgresql
```

- Download [PGAdmin](https://www.pgadmin.org/) client to interact with Postgres Data-base.

- Install [sqlc](https://sqlc.dev/): SQLC is an amazing Go program that generates Go code from SQL queries. It's not exactly an ORM, but rather a tool that makes working with raw SQL almost as easy as using an ORM.
```
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

```

- Configure *sqlc*:
```
version: "2"
sql:
  - schema: "sql/schema"
    queries: "sql/queries"
    engine: "postgresql"
    gen:
      go:
        out: "internal/database"
```

- Write an SQL query to be created as GO method/function:
  Inside the sql/queries directory, create a file called users.sql. Here is mine:
```
-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES ($1, $2, $3, $4)
RETURNING *;
```
$1, $2, $3, and $4 are parameters that we'll be able to pass into the query in our Go code. The `:one` at the end of the query name tells SQLC that we expect to get back a single row (the created user).
Other keywords are `:many`, `:exec`(execute with void)

- Run `sqlc generate` to generate new Go code for your queries.

- Install [goose](https://github.com/pressly/goose): Goose is a database migration tool written in Go. It runs migrations from the same SQL files that SQLC uses, making the pair of tools a perfect fit.
```
go install github.com/pressly/goose/v3/cmd/goose@latest

```

A **migration** is a SQL file that describes a change to your database schema. For now, we need our first migration to create a users table. The simplest format for these files is:

*number*_*name*.sql
For example, I created a file in sql/schema called 001_users.sql with the following contents:

```
-- +goose Up
CREATE TABLE ...

-- +goose Down
DROP TABLE users;
```

To run the *migration*:
```
goose postgres CONN up
ex: goose postgres postgres://<<username>>:<<password>>@<<address>>:<<port>>/<<db-name>> up
```
