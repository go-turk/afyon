run-test:
	docker run --name acPostgres -e POSTGRES_PASSWORD=acPass123 -d -p 5432:5432  postgres
	sleep 1
	- go run cmd/api/main.go
	docker stop acPostgres
	docker rm acPostgres
