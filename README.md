# football-news-backend

Pet project to daily fetch personalized news about football

## Running locally

### Local

Export env variables. Follow `examples.env` file to export required variables.

Start a `football` postgres database, for example, with Docker:

```
docker run \
--rm --name postgres \
-e POSTGRES_USER=postgres \
-e POSTGRES_PASSWORD=postgres \
-e POSTGRES_DB=football \
-p 5432:5432 \
-d postgres:latest
```

Run the app with `go run main.go`

### With docker-compose

Build app and db containers with `docker-compose up -d`

## Port

- 8080

## Endpoints

- `/v1/news` - to get latest news about football
- `/v1/ask` - to ask about football things
