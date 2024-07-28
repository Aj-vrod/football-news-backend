FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./ ./

ARG dns="postgres://postgres:postgres@db:5432/football?sslmode=disable"
ARG migrations_path="file://migrations"
ARG gemini_key="dummy"

ENV POSTGRESQL_DNS=${dns}
ENV MIGRATIONS_PATH=${migrations_path}
ENV GEMINI_KEY=${gemini_key}

RUN CGO_ENABLED=0 GOOS=linux go build -o /football-news

EXPOSE 8080

CMD ["/football-news"]
