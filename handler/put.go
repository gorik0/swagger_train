package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"swagger/data"
)


// swagger:route PUT /{id} products updateProduct
// Update product with id
//responses:
//200: noContentResponse
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
