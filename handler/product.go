package handler

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"swagger/data"
)

type ProductsHandler struct {
	l *log.Logger
}

func (h ProductsHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	h.l.Println("GetProducts")
	products := data.GetProducts()
	err := products.ToJSON(w)
	if err != nil {
		http.Error(w, "error while marshalling", http.StatusInternalServerError)
		return
	}
}

func (h ProductsHandler) PostProducts(w http.ResponseWriter, r *http.Request) {
	h.l.Println("PostProducts")

	h.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

type KeyProduct struct{}

func (h ProductsHandler) PutProducts(w http.ResponseWriter, r *http.Request) {
	h.l.Println("PutProducts")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.PutProducts(id, &prod)
	if err == data.ErrNoProductFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}

}

func (h ProductsHandler) MiddlewareProductsValidate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var prod *data.Product

		err := prod.FromJSON(r.Body)
		if err != nil {

			http.Error(w, fmt.Sprintf("Error while unmarshalling product ::: %s", err), http.StatusBadRequest)
			return
		}

		err = prod.Validate()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error while validating product :::: %s", err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)

	})
}

func NewProductHandler(l *log.Logger) *ProductsHandler {
	return &ProductsHandler{l: l}
}