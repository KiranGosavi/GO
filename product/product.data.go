package product

import (
	"context"
	"database/sql"
	"time"

	"github.com/KiranGosavi/webservice/database"
)


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

func removeProduct(productID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, err := database.DbConn.ExecContext(ctx, `DELETE FROM products where productId=?`, productID)
	if err != nil {
		return err
	}
	return nil
}


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
