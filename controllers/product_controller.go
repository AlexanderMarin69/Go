package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// RegisterProductRoutes registers all product-related routes for v1 API
func RegisterProductRoutes(r chi.Router) {
	r.Route("/products", func(r chi.Router) {
		r.Get("/", listProducts)
		r.Post("/", createProduct)
		r.Get("/{id}", getProduct)
		r.Put("/{id}", updateProduct)
		r.Delete("/{id}", deleteProduct)
	})
}

// listProducts returns a list of all products
func listProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"list all products"}`))
}

// createProduct creates a new product
func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message":"product created"}`))
}

// getProduct retrieves a product by ID
func getProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"get product","id":"` + id + `"}`))
}

// updateProduct updates a product by ID
func updateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"product updated","id":"` + id + `"}`))
}

// deleteProduct deletes a product by ID - returns bad request for testing
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(`{"error":"cannot delete product at this time","code":"OPERATION_NOT_ALLOWED"}`))
}
