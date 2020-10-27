package product

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/websocket"
)

const productsBasePath = "products"

//const productsBasePath=

//SetupRoutes
func SetupRoutes(apiBasePath string) {
	handleProducts := http.HandlerFunc(productHandler)
	handleProduct := http.HandlerFunc(productHandlerSingle)
	http.Handle("/websocket", websocket.Handler(productSocket))
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, productsBasePath), handleProducts)
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, productsBasePath), handleProduct)
}

func productHandlerSingle(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, "products/")
	//to fect last argument of url
	productID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	product, err := getProduct(productID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		//return single product
		productJSON, err := json.Marshal(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productJSON)
	case http.MethodPut:
		//update product in the list
		var product Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if product.ProductID != productID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = updateProduct(product)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	case http.MethodOptions:
		return
	case http.MethodDelete:
		removeProduct(productID)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return

	}
}
func productHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		productList, err := getProductList()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		productJSON, err := json.Marshal(productList)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(productJSON)
		if err != nil {
			log.Fatal(err)
		}

	case http.MethodPost:
		var product Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = insertProduct(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		return
	case http.MethodOptions:
		return
	}
}
