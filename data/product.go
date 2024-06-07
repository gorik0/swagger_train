package data

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"io"
	"regexp"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func (p *Product) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)

}

func (p *Product) Validate() error {

	v := validator.New()
	v.RegisterValidation("sku", validateSKU)
	return v.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)
	if len(matches) != 1 {
		return false
	}

	return true
}

type Products []*Product

// https://golang.org/pkg/encoding/json/#NewEncoder
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}
func AddProduct(p *Product) {
	p.ID = nextId()
	productList = append(productList, p)
}

func nextId() int {
	lastProdID := productList[len(productList)-1].ID
	return lastProdID + 1
}

var ErrNoProductFound = fmt.Errorf("No product found!!!")

func PutProducts(id int, p *Product) error {
	pos := productPos(id)
	if pos == -1 {
		return ErrNoProductFound
	}
	productList[pos] = p
	return nil

}

func productPos(id int) int {
	for i, product := range productList {
		if product.ID == id {
			return i
		}

	}
	return -1
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
