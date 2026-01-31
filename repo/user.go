package repo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	IsAdmin  bool   `json:"is_admin" db:"is_admin"`
}

type UserRepo interface {
	Create(u User) (*User, error)
	Get(email string, password string) (*User, error)
	// 	List() ([]*User, error)
	// 	Update(updatedUser User) (*User, error)
	// 	Delete(UserId int) error
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(newUser User) (*User, error) {

	query := `
		INSERT INTO users (
			username,
			email, 
			password,
			is_admin
		)

		VALUES (
			:username,
			:email, 
			:password,
			:is_admin
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

func (r *userRepo) Get(email string, password string) (*User, error) {
	var user User
	query := `
		SELECT id, username, email, password, is_admin
		FROM users
		WHERE email = $1 AND password = $2
		LIMIT 1
	`
	err := r.db.Get(&user, query, email, password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// func (r *userRepo) List() ([]*User, error) {

// 	return r.userList, nil
// }

// func (r *userRepo) Update(updatedUser User) (*User, error) {
// 	for idx, user := range r.userList {
// 		if user.ID == updatedUser.ID {
// 			r.userList[idx] = updatedUser
// 			return &updatedUser, nil
// 		}
// 	}
// 	return nil, nil
// }

// func (r *userRepo) Delete(userId int) error {
// 	for idx, user := range r.userList {
// 		if user.ID == userId {
// 			r.userList = append(r.userList[:idx], r.userList[idx+1:]...)
// 		}
// 	}
// 	return nil
// }
