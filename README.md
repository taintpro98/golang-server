curl http://localhost:3000/health

### Docker
- Xoa cac container khong trong file docker-compose.dev.yml
  ```
  docker-compose -f docker-compose.dev.yml up --build -d --remove-orphans 
  ```

### Database
- go run ./cmd/migration/main.go -dir migrations create ${FILE_NAME} sql
- go run ./cmd/migration/main.go -dir migrations up