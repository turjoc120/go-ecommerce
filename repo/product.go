package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Price int    `json:"price" db:"price"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	List() ([]*Product, error)
	Get(productId int) (*Product, error)
	Update(updatedProduct Product) (*Product, error)
	Delete(productId int) error
}

type productRepo struct {
	db *sqlx.DB
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

func (r *productRepo) Get(id int) (*Product, error) {
	var product Product
	query := `
	SELECT 
		id,
		name,
		price
	FROM products
	WHERE id = $1
	`
	err := r.db.Get(&product, query, id)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepo) List() ([]*Product, error) {
	var productList []*Product
	query := `
	SELECT id, name, price
	from  products
	`
	err := r.db.Select(&productList, query)
	if err != nil {
		return nil, err
	}
	return productList, nil
}

func (r *productRepo) Update(p Product) (*Product, error) {
	query := `
		UPDATE products
		SET name=$1, price=$2
		WHERE id = $3
		RETURNING id, name, price
	`
	err := r.db.Get(&p, query, p.Name, p.Price, p.ID)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *productRepo) Delete(id int) error {
	query := `
		DELETE FROM products WHERE id = $1
	`

	res, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	//how many rows he deleted
	rows, _ := res.RowsAffected()
	if rows == 0 {
		//if none
		return sql.ErrNoRows
	}
	//no error, product delted
	return nil
}
