docker-db:
	docker-compose down && docker-compose up -d

run:
	go run main.go
	