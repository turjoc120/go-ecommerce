package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID    int    `json:"id" db:"id"`
	Name  string `json:"name" name:"name"`
	Price int    `json:"price" price:"price"`
}

type productRepo struct {
	db *sqlx.DB
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productId int) (*Product, error)
	List() ([]*Product, error)
	Delete(productId int) error
	Update(p Product) (*Product, error)
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(p Product) (*Product, error) {
	query := `
	INSERT INTO products(
		name,
		price
	) VALUES(
		$1,
		$2
	)
	RETURNING id
	`
	row := r.db.QueryRow(query, p.Name, p.Price)
	err := row.Scan(&p.ID)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *productRepo) Get(productId int) (*Product, error) {
	var prd Product
	query := `
SELECT 
id, 
name, 
price
from products
where id=$1
`
	err := r.db.Get(&prd, query, productId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &prd, nil
}

func (r *productRepo) List() ([]*Product, error) {
	var products []*Product
	query := `
	SELECT 
	id, 
	name, 
	price
	from products
	`
	err := r.db.Select(&products, query)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepo) Delete(productId int) error {
	query := `
	DELETE FROM products
	WHERE id=$1`
	_, err := r.db.Exec(query, productId)
	if err != nil {
		return err
	}
	return nil
}

func (r *productRepo) Update(product Product) (*Product, error) {
	query := `
	UPDATE products
	SET name=$1,
	price=$2
	WHERE id=$3
	`
	_, err := r.db.Exec(query, product.Name, product.Price, product.ID)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
