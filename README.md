# football-news-backend

Pet project to daily fetch personalized news about football and ask LLM model about football related questions.

Find infrastructure digragram [here](https://miro.com/welcomeonboard/QmdaZmNUV0xBQkQxc1hxYzFGVFpVRk9Jd3B0S1FiNGR1SUlvS0YyeEpqblhmTXdWNHFjVGhQcVZQdnFmT1l1ZnwzNDU4NzY0NTkzODY5NDIzMDQ4fDI=?share_link_id=410034112358)

## Running locally

### Local

1. Export env variables. Follow `examples.env` file to export required variables.

2. Start a `football` postgres database, for example, with Docker:

```
docker run \
--rm --name postgres \
-e POSTGRES_USER=postgres \
-e POSTGRES_PASSWORD=postgres \
-e POSTGRES_DB=football \
-p 5432:5432 \
-d postgres:latest
```

3. Run the app with `go run main.go`

### With docker-compose

1. Build app and db containers with `docker-compose up -d`

## Port

- 8080

## Endpoints

- `/v1/news?query=example-query` - to get latest news about football
- `/v1/ask` - to ask about football things

## Planned tasks (in no specific order)

- [ ] Update Dockerfile to include Gemini token
- [ ] Move Docker env outside, to prevent hardcoded values
- [ ] Create curated prompt for Gemini
- [ ] Create DB client to enable `INSERT` and `SELECT` queries
- [ ] Create Storage package to store data into DB
- [ ] Prevent app to fail when rerunning migration
- [ ] Create job to trigger scrapper and store incoming news
- [ ] Create cronjob to trigger scraper and storage daily at 8am
- [ ] Create reporter to gather news (either from cronjob or by calling the db client) and expose the news in /v1/news
- [ ] Create persisted DB as AWS RDS
- [ ] Run app as EC2 instance
- [ ] Create news filter (pass news that include specific terms)
