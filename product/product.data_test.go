package product

import (
	"reflect"
	"testing"
)

func TestGetTopTenProducts(t *testing.T) {
	tests := []struct {
		name    string
		want    []Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTopTenProducts()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTopTenProducts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTopTenProducts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getProduct(t *testing.T) {
	type args struct {
		productID int
	}
	tests := []struct {
		name    string
		args    args
		want    *Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getProduct(tt.args.productID)
			if (err != nil) != tt.wantErr {
				t.Errorf("getProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getProduct() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getProductList(t *testing.T) {
	tests := []struct {
		name    string
		want    []Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getProductList()
			if (err != nil) != tt.wantErr {
				t.Errorf("getProductList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getProductList() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insertProduct(t *testing.T) {
	type args struct {
		product Product
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := insertProduct(tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("insertProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("insertProduct() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeProduct(t *testing.T) {
	type args struct {
		productID int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := removeProduct(tt.args.productID); (err != nil) != tt.wantErr {
				t.Errorf("removeProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_updateProduct(t *testing.T) {
	type args struct {
		product Product
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := updateProduct(tt.args.product); (err != nil) != tt.wantErr {
				t.Errorf("updateProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
