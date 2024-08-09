# football-news-backend

Pet project to daily fetch personalized news about football and ask LLM model about football related questions.

Find infrastructure digragram [here](https://miro.com/welcomeonboard/QmdaZmNUV0xBQkQxc1hxYzFGVFpVRk9Jd3B0S1FiNGR1SUlvS0YyeEpqblhmTXdWNHFjVGhQcVZQdnFmT1l1ZnwzNDU4NzY0NTkzODY5NDIzMDQ4fDI=?share_link_id=410034112358)

## Running locally

### For development

1. List the required env variables in a `.env` file. Follow `examples.env` file to learn what they are.

2. Comment out the app service from docker-compose.yaml

3. Start the db with `docker-compose up -d`

4. Run the app with `go run main.go`

### For simple building

1. List the required env variables in a `.env` file. Follow `examples.env` file to learn what they are.

2. Build app and db containers with `docker-compose up -d`

## Port

- 8080

## Endpoints

- `/v1/news` - to get all news about football
- `/v1/news?date=DD.MM.YYYY` - to get news about football of a specific date
- `/v1/ask?query=example-query` - to ask about football things

## Planned tasks (in no particular order)

- [x] Update Dockerfile to include Gemini token
- [x] Move Docker env outside, to prevent hardcoded values
- [ ] Create file document with LLM to collect knowledgebase about football
- [ ] Store the knowledge base in an S3 bucket
- [ ] Retrieve the S3 bucket to serve as context for LLM
- [ ] Create curated prompt for Gemini
- [x] Create validation for "date" and "query" values
- [x] Create DB handler to enable `INSERT` and `SELECT` queries
- [x] Create Storage package to store data into DB
- [ ] Prevent app to fail when rerunning migration
- [ ] Create job to trigger scrapper and store incoming news
- [ ] Create cronjob to trigger scraper and storage daily at 8am
- [ ] Create reporter to gather news (either from cronjob or by calling the db client) and expose the news in /v1/news
- [ ] Create persisted DB as AWS RDS
- [ ] Run app as EC2 instance
- [ ] Create news filter (pass news that include specific terms)
