FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./ ./

ENV POSTGRESQL_DNS="postgres://postgres:password@localhost:5432/football?sslmode=disable"
ENV MIGRATIONS_PATH="file://migrations"
RUN CGO_ENABLED=0 GOOS=linux go build -o /football-news

EXPOSE 8080

CMD ["/football-news"]
