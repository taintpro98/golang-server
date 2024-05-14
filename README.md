curl http://localhost:3000/health

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
- go run ./cmd/migration/main.go -dir migrations create ${FILE_NAME} sql
- go run ./cmd/migration/main.go -dir migrations up