package models

import (
	"reflect"
	"testing"
)

func TestProduct(t *testing.T) {
	product := Product{
		ProductID:        1,
		Name:             "Test Product",
		Description:      "This is a test product.",
		Images:           []string{"image1.jpg", "image2.jpg"},
		Price:            1000,
		CompressedImages: []string{"compressed_image1.jpg", "compressed_image2.jpg"},
	}

	// Test for ProductID
	if product.ProductID != 1 {
		t.Errorf("Expected ProductID to be 1, got %d", product.ProductID)
	}

	// Test for Name
	expectedName := "Test Product"
	if product.Name != expectedName {
		t.Errorf("Expected Name to be %s, got %s", expectedName, product.Name)
	}

	// Test for Description
	expectedDescription := "This is a test product."
	if product.Description != expectedDescription {
		t.Errorf("Expected Description to be %s, got %s", expectedDescription, product.Description)
	}

	// Test for Images
	expectedImages := []string{"image1.jpg", "image2.jpg"}
	if !reflect.DeepEqual(product.Images, expectedImages) {
		t.Errorf("Expected Images to be %v, got %v", expectedImages, product.Images)
	}

	// Test for Price
	expectedPrice := 1000
	if product.Price != expectedPrice {
		t.Errorf("Expected Price to be %d, got %d", expectedPrice, product.Price)
	}

	// Test for CompressedImages
	expectedCompressedImages := []string{"compressed_image1.jpg", "compressed_image2.jpg"}
	if !reflect.DeepEqual(product.CompressedImages, expectedCompressedImages) {
		t.Errorf("Expected CompressedImages to be %v, got %v", expectedCompressedImages, product.CompressedImages)
	}
}

func TestNewProduct(t *testing.T) {
	newProduct := NewProduct{
		UserID:             1,
		ProductName:        "Test Product",
		ProductDescription: "This is a test product.",
		ProductImages:      []string{"image1.jpg", "image2.jpg"},
		ProductPrice:       1000,
	}

	// Test for UserID
	if newProduct.UserID != 1 {
		t.Errorf("Expected UserID to be 1, got %d", newProduct.UserID)
	}

	// Test for ProductName
	expectedProductName := "Test Product"
	if newProduct.ProductName != expectedProductName {
		t.Errorf("Expected ProductName to be %s, got %s", expectedProductName, newProduct.ProductName)
	}

	// Test for ProductDescription
	expectedProductDescription := "This is a test product."
	if newProduct.ProductDescription != expectedProductDescription {
		t.Errorf("Expected ProductDescription to be %s, got %s", expectedProductDescription, newProduct.ProductDescription)
	}

	// Test for ProductImages
	expectedProductImages := []string{"image1.jpg", "image2.jpg"}
	if !reflect.DeepEqual(newProduct.ProductImages, expectedProductImages) {
		t.Errorf("Expected ProductImages to be %v, got %v", expectedProductImages, newProduct.ProductImages)
	}

	// Test for ProductPrice
	expectedProductPrice := 1000
	if newProduct.ProductPrice != expectedProductPrice {
		t.Errorf("Expected ProductPrice to be %d, got %d", expectedProductPrice, newProduct.ProductPrice)
	}
}
