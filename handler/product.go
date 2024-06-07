// Packcage classification of prod API
//
// Documentationfor Product API
// Schemes:http
// BasePath:  /
// Version:1.0.0
// Produces:
//
//	-application/json
//
// Consumes:
//
//	-application/json
//
//swagger:meta
package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"swagger/data"
)

type ProductsHandler struct {
	l *log.Logger
}


type GenericError struct {
	Message string `json:"message"`
}

type ValidationError struct {
	Messages []string `json:"messages"`
}

type KeyProduct struct{}

func (h ProductsHandler) MiddlewareProductsValidate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var prod data.Product

		err := data.FromJSON(&prod, r.Body)
		//de := json.NewDecoder(r.Body)
		//err := de.Decode(prod)
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
