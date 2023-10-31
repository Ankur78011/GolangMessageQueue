package models

type Product struct {
	ProductID        int      `db:"product_id"`
	Name             string   `db:"product_name"`
	Description      string   `db:"product_description"`
	Images           []string `db:"product_images"`
	Price            int      `db:"product_price"`
	CompressedImages []string `db:"compressed_product_images"`
}
