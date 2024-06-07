package handler

import (
	"net/http"
	"swagger/data"
)
// swagger:route get / products noParam
// Returns list of products
//responses:
//200: productResponse
//400: errorValidate
func (h ProductsHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	h.l.Println("GetProducts")
	products := data.GetProducts()
	err := products.ToJSON(w)
	if err != nil {
		http.Error(w, "error while marshalling", http.StatusInternalServerError)
		return
	}
}
