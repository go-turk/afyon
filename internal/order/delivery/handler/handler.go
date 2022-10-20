package handler

import (
	"encoding/json"
	"github.com/go-turk/afyon/internal/order"
	"github.com/go-turk/afyon/internal/order/domain"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	router       *mux.Router
	orderService order.OrderService
}

func NewOrderHandler(router *mux.Router, orderService order.OrderService) (*OrderHandler, error) {
	handler := &OrderHandler{
		router:       router,
		orderService: orderService,
	}
	handler.RegisterRoutes()
	return handler, nil
}

func (h *OrderHandler) RegisterRoutes() {
	h.router.HandleFunc("/orders", h.createOrder).Methods("POST")
	h.router.HandleFunc("/orders/{id}", h.getOrder).Methods("GET")
	h.router.HandleFunc("/orders/{id}", h.updateOrder).Methods("PUT")
	h.router.HandleFunc("/orders/{id}", h.deleteOrder).Methods("DELETE")
}

func (h *OrderHandler) createOrder(w http.ResponseWriter, r *http.Request) {
	var orderRequest struct {
		Items []int `json:"items"`
	}
	json.NewDecoder(r.Body).Decode(&orderRequest)
	items, err := h.orderService.FindItems(orderRequest.Items...)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	order := domain.NewOrder(items, 1)
	err = h.orderService.Create(order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) getOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	order, err := h.orderService.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) updateOrder(w http.ResponseWriter, r *http.Request) {

}

func (h *OrderHandler) deleteOrder(w http.ResponseWriter, r *http.Request) {

}
