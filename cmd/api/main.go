package main

import (
	"fmt"
	"github.com/go-turk/afyon/internal/order/delivery/handler"
	"github.com/go-turk/afyon/internal/order/domain"
	"github.com/go-turk/afyon/internal/order/repository/pg"
	"github.com/go-turk/afyon/internal/order/service"
	"github.com/go-turk/afyon/pkg/db/postgres"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	db, err := postgres.Connect()
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	itemRepo := pg.NewItemRepository(db)
	orderRepo := pg.NewOrderRepository(db)
	orderService := service.NewOrderService(itemRepo, orderRepo)
	ufakBaslangicTesti(orderService)
	_, err = handler.NewOrderHandler(r, orderService)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server started on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}

func ufakBaslangicTesti(orderService *service.OrderService) {
	// Test Çıktıları
	item := domain.NewItem("test", 10.7)
	item2 := domain.NewItem("test2", 109.5)
	orderService.CreateItem(item)
	orderService.CreateItem(item2)
	order := domain.NewOrder([]domain.Item{*item}, 1)
	order2 := domain.NewOrder([]domain.Item{*item, *item2}, 2)
	order3 := domain.NewOrder([]domain.Item{*item, *item2}, 33)
	fmt.Println(orderService.Create(order))
	fmt.Println(orderService.Create(order2))
	fmt.Println(orderService.Create(order3))
	fmt.Println("1 numaralı:")
	fmt.Println(orderService.FindByID(1))
	fmt.Println("2 numaralı:")
	fmt.Println(orderService.FindByID(2))
	fmt.Println("3 numaralı:")
	fmt.Println(orderService.FindByID(3))
	fmt.Println("Tüm Orderlar:")
	fmt.Println(orderService.FindAll())
	fmt.Println("1,2 numaralı:")
	fmt.Println(orderService.FindAll(1, 2))
	// Test Çıktıları
}
