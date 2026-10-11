# Layered Student API with Google Wire

A CRUD REST API in Go that uses **Google Wire** for compile-time dependency injection, with
**Gin** for HTTP and **GORM** on **PostgreSQL**.

The point of this project is the structure: each layer depends only on the one below it, and
Wire connects them.

```
mypackage/
  cmd/          main.go, wire.go (injector), wire_gen.go (generated)
  router/       routes
  controller/   HTTP handlers
  service/      business logic
  repository/   database access
  model/        the Student model
  database/     database connection
```

`database → repository → service → controller → router`

## Endpoints

| Method | Path | Action |
|---|---|---|
| `POST` | `/students` | Create a student |
| `GET` | `/students` | List all students |
| `GET` | `/students/:id` | Get one student |
| `PUT` | `/students/:id` | Update a student |
| `DELETE` | `/students/:id` | Delete a student |

## Run

```sh
cd mypackage
export DATABASE_URL="host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
go run ./cmd
```

The server starts on `:8080`. To regenerate the injector after changing providers:

```sh
go run github.com/google/wire/cmd/wire ./cmd
```
