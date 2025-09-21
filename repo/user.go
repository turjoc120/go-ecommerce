package repo

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

type userRepo struct {
	userList []*User
}

type UserRepo interface {
	Create(user User) (*User, error)
	Get(email string, pass string) (*User, error)
	// List() ([]*User, error)
	// Delete(userId int) error
	// Update(user User) (*User, error)
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (r *userRepo) Create(u User) (*User, error) {
	if u.ID != 0 {
		return &u, nil
	}
	u.ID = len(r.userList) + 1
	r.userList = append(r.userList, &u)
	return &u, nil
}

func (r *userRepo) Get(email string, pass string) (*User, error) {
	for _, u := range r.userList {
		if u.Email == email && u.Password == pass {
			return u, nil
		}
	}
	return nil, nil
}
