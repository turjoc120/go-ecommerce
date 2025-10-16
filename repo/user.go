package repo

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          int    `json:"id" db:"id"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	IsShopOwner bool   `json:"is_shop_owner" db:"is_shop_owner"`
}

type userRepo struct {
	db *sqlx.DB
}

type UserRepo interface {
	Create(user User) (*User, error)
	Find(email string, pass string) (*User, error)
	// List() ([]*User, error)
	// Delete(userId int) error
	// Update(user User) (*User, error)
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(newUser User) (*User, error) {
	query := `INSERT INTO users(
first_name,
last_name,
email,
password,
is_shop_owner
)

VALUES(
:first_name,
:last_name,
:email,
:password,
:is_shop_owner
)
RETURNING id
`

	var userId int
	rows, err := r.db.NamedQuery(query, newUser)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if rows.Next() {
		rows.Scan(&userId)
	}
	newUser.ID = userId
	return &newUser, nil
}

func (r *userRepo) Find(email string, pass string) (*User, error) {
	var user User
	query := `
	SELECT id, first_name, last_name, email, password, is_shop_owner 
	FROM users 
	WHERE email = $1 AND password = $2
	LIMIT 1
	`

	err := r.db.Get(&user, query, email, pass)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil

}
