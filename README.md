curl http://localhost:3000/health

### Docker
- docker-compose -f docker-compose.dev.yml up --build -d --remove-orphans

### Database
- go run ./cmd/migration/main.go -dir migrations up