package handler

import (
	"net/http"
	"swagger/data"
)
// swagger:route POST / products createProduct
// Create product
//responses:
//200: productResponse

func (h ProductsHandler) PostProducts(w http.ResponseWriter, r *http.Request) {
	h.l.Println("PostProducts")


	h.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
	w.WriteHeader(http.StatusNoContent)
}
