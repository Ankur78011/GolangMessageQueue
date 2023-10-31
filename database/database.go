package database

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
	"zocket.example.com/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "Zocket"
)

type Database struct {
	Db *sql.DB
}

func NewDatabase() (*Database, error) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return nil, err
	}
	fmt.Println("database connection successful")
	return &Database{Db: db}, nil
}

// Sending Prodcut details
func (dbInstance *Database) StoreProduct(newproduct *models.NewProduct) (int64, error) {

	sqlStatement := `INSERT INTO products (product_name, product_description, product_images, product_price) VALUES ($1, $2, $3, $4) RETURNING product_id`

	var productID int64

	err := dbInstance.Db.QueryRow(sqlStatement, newproduct.ProductName, newproduct.ProductDescription, pq.Array(newproduct.ProductImages), newproduct.ProductPrice).Scan(&productID)
	if err != nil {
		fmt.Println("Cannot insert new record because of", err)
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
		fmt.Println("cannot insert local file path in database")
		return err
	}
	return nil
}
