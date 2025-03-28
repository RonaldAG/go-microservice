package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Product defines the structure of an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Product) Validate() error {
	validator := validator.New()
	validator.RegisterValidation("sku", validateSKU)
	return validator.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	// sku is format abc-adf-adsfg
	reg := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]`)
	matches :=	reg.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}
	return true
}


type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, prodId, err := findProduct(id)

	if err != nil {
		return err
	}

	p.ID = prodId
	productList[prodId - 1] = p
	return nil
}

func DeleteProduct(id int) error {
	_, prodId, err := findProduct(id)
	index := prodId - 1
	if err != nil {
		return err
	}

	productList = append(productList[:index], productList[index+1:]...)
	return nil
}

var ErrorProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for _, p := range productList {
		if p.ID == id {
			return p, id, nil
		}
	}
	return nil, -1, ErrorProductNotFound
}

func getNextID() int {
	lp := productList[len(productList) - 1]
	return lp.ID + 1
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milk coffe",
		Price:       2.45,
		SKU:         "abc123",
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffe without milk",
		Price:       1.99,
		SKU:         "fjd34",
	},
}
