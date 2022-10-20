##Endpoints

- GET /orders/:id (1,2,3 id'leri mevcuttur)
- POST /orders (örnek request body: `{"id":[1,2,3]}`)


##Notes
- mkdir -p {cmd/api,internal/{user,order}/{domain,service,repository,delivery/handler},pkg/db} >> Create Folders
docker run --name acPostgres -e POSTGRES_PASSWORD=acPass123 -d -p 5432:5432 -v ./docker:/var/lib/postgresql/data  postgres
