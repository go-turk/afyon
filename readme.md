# Golang ile Server oluşturma (4. Hafta)
## Proje Başlığı: Afyon Yöresel Yemekleri Backend
### Endpointler

- /yemek {GET} : Tüm yemekleri getirir.
- /yemek/{id} {GET} : İd'si verilen yemeği getirir.
- /yemek/{id} {DELETE} : İd'si verilen yemeği siler.
- /yemek/{id} {PUT} : İd'si verilen yemeği günceller.
- /yemek {POST} : Yeni bir yemek ekler.

### Kullanılan Kütüphaneler
- github.com/gorilla/mux (Router)
- github.com/jinzhu/gorm (ORM)
- github.com/joho/godotenv (Env Dosyası)

### Ogrenilecekler
- Golang ile Server oluşturma
- Golang ile CRUD işlemleri


##Notes
- mkdir -p {cmd/api,internal/{user,order}/{domain,service,repository,delivery/handler},pkg/db} >> Create Folders
docker run --name acPostgres -e POSTGRES_PASSWORD=acPass123 -d -p 5432:5432 -v ./docker:/var/lib/postgresql/data  postgres
