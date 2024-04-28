# Description

Student API

## Folder Structure

```
└── student_api
    └── api.http
    └── cmd
        └── api
            └── db.go
            └── handlers.go
            └── main.go
            └── middlewares.go
            └── routes.go
            └── utils.go
    └── go.mod
    └── go.sum
    └── internal
        └── models
            └── StudentModel.go
        └── repository
            └── dbrepo
                └── MySQL_dbrepo.go
            └── repository.go
    └── sql
        └── db.sql
```

## How to run localy

Clone the project to your machine

```bash
cd student_api
```

Run `db.sql` file in the /sql folder in MySQL to initialize the database on localhost.

In the project directory, you can run:

Install the mysql driver and chi framework.

```bash
go get -u github.com/go-sql-driver/mysql
go get -u github.com/go-chi/chi/v5
```

Run the server: provide the appropriate information to connect to the database, `env` can be `dev` or `prod`

```bash
go run ./cmd/api -env= -DBUsername= -DBPassword= -DBHost= -DBPort=
```
