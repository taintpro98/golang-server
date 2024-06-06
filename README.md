### Features
- Authenticaion using JWT
- Realtime reserving seats concurrently using Redis, Postgresql
- Realtime searching a lot of records using Elastic Search
- Realtime newsfeed using SSE, Asynq (or Kafka)
- Peer to peer messages using Websocket

### Technologies
- Elastic Search
- Big Query
- Kafka
### Ingredients
- API and SSE server using Gin
- Event dispatcher consuming Kafka Events
- Cron Jobs
- Workers using Asynq with Redis
- Websocket server using ...
### Health check
```
curl http://localhost:5000/health
```
### Docker
- Remove containers not in the file docker-compose.dev.yml
  ```
  docker-compose -f docker-compose.dev.yml up --build -d --remove-orphans 
  ```
- Run kafka containers
  ```
  docker-compose -f docker-compose.kafka.yml up --build -d
  ```
- Run elastic containers
  ```
  docker-compose -f docker-compose.elastic.yml up --build -d
  ```

### Database
#### Migrations
```
go run ./cmd/migration/main.go -dir migrations create ${FILE_NAME} sql
go run ./cmd/migration/main.go -dir migrations up
```
#### Seeds
```
go run ./cmd/migration/main.go -dir seeds create ${FILE_NAME} sql
go run ./cmd/migration/main.go -dir seeds up
```

### Elastic search
- [Dashboard](http://localhost:5601/)
```
{
  "query": {
    "match_all": {}
  }
}
{
  "query": {
    "wildcard": {
      "phone": "*191954*"
    }
  }
}
{
    "query": {
        "bool": {
            "should": [
                {
                    "wildcard": {
                        "phone": "*191*"
                    }
                },
                {
                    "wildcard": {
                        "email": "*191*"
                    }
                }
            ]
        }
    }
}
```
### Loadtest
```
k6 run k6/loadtest.js
```
### Server Sent Events
- We can use Asynq or Kafka instead of channel to handle Server Sent Events, Kafka for microservices architecture ?
- Because each user needs a particular channel to consume others post ? 

### References
- [Streaming Server-Sent Events With Go](https://pascalallen.medium.com/streaming-server-sent-events-with-go-8cc1f615d561)