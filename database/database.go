package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
	"zocket.example.com/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
)

type Database struct {
	Db *sql.DB
}

func NewDatabase() (*Database, error) {
	// for makeing a databse when you run this file
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`create database zocket`)
	if err != nil {
		log.Printf("cannot connect to database")
	}
	//accessing database

	psqlconntwo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=zocket sslmode=disable", host, port, user, password)

	dbtwo, err := sql.Open("postgres", psqlconntwo)

	_, err = dbtwo.Exec(`CREATE TABLE users IF NOT EXISTS (
		id serial PRIMARY KEY,
		name  text,
		mobile text,
		latitude DOUBLE PRECISION,
		longitude DOUBLE PRECISION,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`)
	if err != nil {
		log.Println("Cannot create user table", err)
	}

	_, err = dbtwo.Exec(`CREATE TABLE products IF NOT EXISTS(
		product_id SERIAL PRIMARY KEY,
		product_name text,
		product_description TEXT,
		product_images TEXT[], 
		product_price int,
		compressed_product_images TEXT[], 
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`)
	if err != nil {
		log.Println("Cannot create product table", err)
	}

	log.Println("database connection successful")
	return &Database{Db: dbtwo}, nil
}

// Sending Prodcut details
func (dbInstance *Database) StoreProduct(newproduct *models.NewProduct) (int64, error) {

	sqlStatement := `INSERT INTO products (product_name, product_description, product_images, product_price) VALUES ($1, $2, $3, $4) RETURNING product_id`

	var productID int64

	err := dbInstance.Db.QueryRow(sqlStatement, newproduct.ProductName, newproduct.ProductDescription, pq.Array(newproduct.ProductImages), newproduct.ProductPrice).Scan(&productID)
	if err != nil {
		log.Println("Cannot insert new record because of", err)
		return 0, err
	}

	return productID, nil
}

// Fetching Iamges Url
func (dbInstance *Database) GetProductFromProductId(product_id int) ([]string, error) {
	var currentProductImages pq.StringArray
	sqlstatment := `select product_images from products where product_id=$1`
	err := dbInstance.Db.QueryRow(sqlstatment, product_id).Scan(&currentProductImages)
	if err != nil {
		return nil, err
	}
	return currentProductImages, nil
}

// Adding compressed path
func (dbInstance *Database) AddCompressedProductImages(product_id int, LocalFileName []string) error {
	sql := `update products set compressed_product_images=$1 where product_id=$2`
	psqlArrayFileName := pq.Array(LocalFileName)
	_, err := dbInstance.Db.Exec(sql, psqlArrayFileName, product_id)
	if err != nil {
		log.Println("cannot insert local file path in database")
		return err
	}
	return nil
}
