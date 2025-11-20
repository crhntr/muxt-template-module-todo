# A Simple MUXT Project Template using HTMX

This is an example [muxt](http://github.com/typelate/muxt) hypertext server.
See this article for how to use this template: https://go.dev/blog/gonew.

It is based on https://github.com/crhntr/muxt-template-module-htmx.
You may also want to see:
- [HTMX](https://htmx.org): configured for site interactivity
- [Pico CSS](https://picocss.com/docs): configured to make the site look a bit better by default
- [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter): used to generate fake Server implementations for testing templates
- [unpkg.com](https://unpkg.com): use the meta query param to get new version integrity values

It extends the module by using sqlc to for database interactions with PostgreSQL.
- [sqlc](https://docs.sqlc.dev/en/latest/): to generate type safe database interactions
- [pqx](https://pkg.go.dev/github.com/jackc/pgx/v5): for a database driver
- [goose](https://github.com/pressly/goose): for migrations

Tool versions are managed with [Mise](https://mise.jdx.dev/).

## Developing

Add new html in `./internal/hypertext/templates`.
Use inline template declarations in the template files using muxt template name semantics
then run Go generate to generate a function that maps hypertext concepts to helpful handler methods.

## Testing

After developing routes, run `go generate ./...`.

Write handler tests in `./internal/hypertext/server_test.go`.

Write template tests in `./internal/hypertext/template_routes_test.go`.

Run tests with: `go test ./...`

## Database

This project uses PostgreSQL with pgx/v5.

### Quick Start

```bash
# Start PostgreSQL with Docker
docker compose up -d

# Load environment variables (requires direnv)
direnv allow

# Run the server
go run ./cmd/server
```

### Configuration

The database connection is configured via the `DATABASE_URL` environment variable.

The included `.envrc` sets this automatically when using [direnv](https://direnv.net/), reading credentials directly from `docker-compose.yml`.

### Docker Compose

The `docker-compose.yml` provides a PostgreSQL 17 instance with:
- Database: `todo`
- User: `todo`
- Password: `todo`
- Port: `5432`

```bash
# Start database
docker compose up -d

# Stop database
docker compose down

# Stop and remove data
docker compose down -v
```
