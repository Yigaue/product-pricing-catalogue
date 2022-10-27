package controllers

import (
	"crud-rest-api/database"
	"crud-rest-api/entities"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product entities.Product
	json.NewDecoder(r.Body).Decode(&product)
	database.Instance.Create(&product)
	json.NewEncoder(w).Encode(product)
}

func checkIfProductExist(productId string) bool  {
	var product entities.Product
	database.Instance.First(&product, productId)
	if product.ID == 0 {
		return false
	} else {
		return true
	}
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"] //Get the route var from the request in this Product Id or we can say
	// Gets the Product Id from the Query string of the request.
	if(checkIfProductExist(productId) == false) {
		json.NewEncoder(w).Encode("Product not found!")
		return
	}
	var product entities.Product
	database.Instance.First(&product, productId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}