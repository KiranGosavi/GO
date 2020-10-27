package product

import (
	"context"
	"database/sql"
	"time"

	"github.com/KiranGosavi/webservice/database"
)

// var productMap = struct {
// 	sync.RWMutex
// 	m map[int]Product
// }{m: make(map[int]Product)}

// func init() {
// 	fmt.Println("loading products >>>>>>>>")
// 	prodMap, err := loadProductMap()
// 	productMap.m = prodMap
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%d products loaded...\n", len(productMap.m))
// }

// func loadProductMap() (map[int]Product, error) {
// 	filename := "products.json"
// 	_, err := os.Stat(filename)
// 	if os.IsNotExist(err) {
// 		return nil, fmt.Errorf("file[%s] does not exists", filename)
// 	}

// 	file, _ := ioutil.ReadFile(filename)
// 	productList := make([]Product, 0)
// 	err = json.Unmarshal([]byte(file), &productList)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	prodMap := make(map[int]Product)
// 	for i := 0; i < len(productList); i++ {
// 		prodMap[productList[i].ProductID] = productList[i]
// 	}
// 	return prodMap, nil
// }

// func getProduct(productID int) *Product {
// 	productMap.RLock()
// 	defer productMap.RUnlock()
// 	if product, ok := productMap.m[productID]; ok {
// 		return &product
// 	}
// 	return nil
// }
func getProduct(productID int) (*Product, error) {
	//fmt.Println(productID)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	row := database.DbConn.QueryRowContext(ctx, `SELECT productId,
	manufacturer, 
	sku, 
	upc,
	pricePerUnit,
	quantityOnHand,
	productName
	FROM products
	WHERE productId= ?`, productID)
	//fmt.Println(row)
	product := &Product{}
	err := row.Scan(&product.ProductID,
		&product.Manufacturer,
		&product.Sku,
		&product.Upc,
		&product.PricePerUnit,
		&product.QuantityOnHand,
		&product.ProductName)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return product, nil
}

// func getProductList() []Product {
// 	productMap.RLock()
// 	products := make([]Product, 0, len(productMap.m))
// 	for _, value := range productMap.m {
// 		products = append(products, value)
// 	}
// 	productMap.RUnlock()
// 	return products
// }

//[{"productId":1,"manufacturer":"Johns-Jenkins","sku":"p5z343vdS","upc":"939581000000","pricePerUnit":"497.45","quantityOnHand":9703,"productName":"sticky note"}
func getProductList() ([]Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	results, err := database.DbConn.QueryContext(ctx, `SELECT productId,
	manufacturer, 
	sku, 
	upc,
	pricePerUnit,
	quantityOnHand,
	productName
	FROM products`)
	if err != nil {
		//	fmt.Println("Here")
		return nil, err
	}
	defer results.Close()
	products := make([]Product, 0)
	for results.Next() {
		var product Product
		results.Scan(&product.ProductID,
			&product.Manufacturer,
			&product.Sku,
			&product.Upc,
			&product.PricePerUnit,
			&product.QuantityOnHand,
			&product.ProductName)
		products = append(products, product)
	}
	//fmt.Println(products[0])
	return products, nil
}

//GetTopTenProducts function
func GetTopTenProducts() ([]Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	results, err := database.DbConn.QueryContext(ctx, `SELECT productId,
	manufacturer, 
	sku, 
	upc,
	pricePerUnit,
	quantityOnHand,
	productName
	FROM products ORDER BY quantityOnHand DESC LIMIT 10`)
	if err != nil {
		//	fmt.Println("Here")
		return nil, err
	}
	defer results.Close()
	products := make([]Product, 0)
	for results.Next() {
		var product Product
		results.Scan(&product.ProductID,
			&product.Manufacturer,
			&product.Sku,
			&product.Upc,
			&product.PricePerUnit,
			&product.QuantityOnHand,
			&product.ProductName)
		products = append(products, product)
	}
	//fmt.Println(products[0])
	return products, nil
}

// func removeProduct(productID int) {
// 	productMap.Lock()
// 	defer productMap.Unlock()
// 	delete(productMap.m, productID)
// }
func removeProduct(productID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, err := database.DbConn.ExecContext(ctx, `DELETE FROM products where productId=?`, productID)
	if err != nil {
		return err
	}
	return nil
}

// func getProductIDs() []int {
// 	productMap.RLock()
// 	productIDs := []int{}
// 	for key := range productMap.m {
// 		productIDs = append(productIDs, key)
// 	}
// 	productMap.RUnlock()
// 	sort.Ints(productIDs)
// 	return productIDs
// }

// func getNextProductID() int {
// 	productIDS := getProductIDs()
// 	return productIDS[len(productIDS)-1] + 1
// }

// func addOrUpdateProduct(product Product) (int, error) {
// 	addOrUpdateID := -1
// 	if product.ProductID > 0 {
// 		oldProduct, err := getProduct(product.ProductID)
// 		if err != nil {
// 			return addOrUpdateID, err
// 		}
// 		if oldProduct == nil {
// 			return 0, fmt.Errorf("product id %g doen not exist", product.ProductID)
// 		}
// 		addOrUpdateID = product.ProductID
// 	} else {
// 		addOrUpdateID = getNextProductID()
// 		product.ProductID = addOrUpdateID
// 	}
// 	productMap.Lock()
// 	productMap.m[addOrUpdateID] = product
// 	productMap.Unlock()
// 	return addOrUpdateID, nil

// }
func updateProduct(product Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := database.DbConn.ExecContext(ctx, `UPDATE products SET
	manufacturer=?, 
	sku=?, 
	upc=?,
	pricePerUnit=?,
	quantityOnHand=?,
	productName=?
	WHERE productID=?`, product.Manufacturer,
		product.Sku, product.Upc, product.PricePerUnit, product.QuantityOnHand,
		product.ProductName, product.ProductID)
	if err != nil {
		return err
	}
	return nil
}

func insertProduct(product Product) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := database.DbConn.ExecContext(ctx, `UPDATE products SET
	manufacturer=?, 
	sku=?, 
	upc=?,
	pricePerUnit=?,
	quantityOnHand=?,
	productName=?
	WHERE productID=?`, product.Manufacturer,
		product.Sku, product.Upc, product.PricePerUnit, product.QuantityOnHand,
		product.ProductName, product.ProductID)
	if err != nil {
		return 0, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return int(insertID), nil
}
