package database

import (
	"testing"

	"zocket.example.com/models"
)

// Test case: Database
func TestNewDatabase(t *testing.T) {
	db, err := NewDatabase()
	if err != nil {
		t.Fatalf("Failed to create a new database instance: %v", err)
	}

	if db == nil {
		t.Error("Database instance is nil")
	}
}

// Test case: StoreProduct method
func TestDatabase_StoreProduct(t *testing.T) {
	db, err := NewDatabase()
	if err != nil {
		t.Fatalf("Failed to create a new database instance: %v", err)
	}
	defer db.Db.Close()

	newProduct := &models.NewProduct{
		ProductName:        "Test Product",
		ProductDescription: "Test Description",
		ProductImages:      []string{"image1.jpg", "image2.jpg"},
		ProductPrice:       99,
	}

	productID, err := db.StoreProduct(newProduct)
	if err != nil {
		t.Errorf("Failed to store a new product: %v", err)
	}

	if productID <= 0 {
		t.Errorf("Invalid product ID: %v", productID)
	}
}

// Test case: AddCompressedProductImages method
func TestDatabase_AddCompressedProductImages(t *testing.T) {
	db, err := NewDatabase()
	if err != nil {
		t.Fatalf("Failed to create a new database instance: %v", err)
	}
	defer db.Db.Close()

	productID := 79
	localFilePaths := []string{"path1.gz", "path2.gz"}

	err = db.AddCompressedProductImages(productID, localFilePaths)
	if err != nil {
		t.Errorf("Failed to add compressed product images: %v", err)
	}
}
