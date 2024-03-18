package data

import (
	"encoding/json"
	"errors"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	SKU         string  `json:"sku"`
	Price       float32 `json:"price"`
	Created_on  string  `json:"createdOn"`
	Deleted_on  string  `json:"deletedOn"`
	Updated_on  string  `json:"updatedOn"`
}

type Products []*Product

func (p *Product) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(p)
}

func (p *Products) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(p)
}

func GetProducts() Products {
	return productList
}

func GetProductByID(id int) (*Product, error) {
	for _, product := range productList {
		if product.ID == id {
			return product, nil
		}
	}
	return nil, errors.New("product not found")
}

func AddProduct(p *Product) {
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	product, err := GetProductByID(id)
	if err != nil {
		return err
	}
	p.ID = id
	p.Created_on = product.Created_on
	p.Updated_on = time.Now().UTC().String()
	for index, product := range productList {
		if product.ID == id {
			// Remove the product
			productList = append(productList[:index], productList[index+1:]...)
		}
	}
	productList = append(productList, p)
	return nil
}

func DeleteProductByID(id int) {
	for index, product := range productList {
		if product.ID == id {
			// Remove the product
			productList = append(productList[:index], productList[index+1:]...)
		}
	}
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "latte",
		Price:       2.43,
		SKU:         "abc123",
		Description: "COffee with forthy and milk",
		Created_on:  time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "tea",
		Price:       1.43,
		SKU:         "fdj321",
		Description: "COffee with milk",
		Created_on:  time.Now().UTC().String(),
	},
}
