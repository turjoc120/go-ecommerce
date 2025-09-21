package repo

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type productRepo struct {
	productList []*Product
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productId int) (*Product, error)
	List() ([]*Product, error)
	Delete(productId int) error
	Update(p Product) (*Product, error)
}

func NewProductRepo() ProductRepo {
	repo := productRepo{}
	genInitProduct(&repo)
	return &repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}

func (r *productRepo) Get(productId int) (*Product, error) {
	for _, p := range r.productList {
		if p.ID == productId {
			return p, nil
		}
	}
	return nil, nil
}

func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}

func (r *productRepo) Delete(productId int) error {
	for idx, p := range r.productList {
		if p.ID == productId {
			r.productList = append(r.productList[:idx], r.productList[idx+1:]...)
		}
	}
	return nil
}

func (r *productRepo) Update(product Product) (*Product, error) {
	for idx, p := range r.productList {
		if p.ID == product.ID {
			r.productList[idx] = &product
		}
	}
	return &product, nil
}

func genInitProduct(r *productRepo) {
	pd1 := Product{ID: 1, Name: "Laptop", Price: 1000}
	pd2 := Product{ID: 2, Name: "Laptop", Price: 1000}
	pd3 := Product{ID: 3, Name: "Laptop", Price: 1000}
	pd4 := Product{ID: 4, Name: "Laptop", Price: 1000}
	pd5 := Product{ID: 5, Name: "Laptop", Price: 1000}

	r.productList = append(r.productList, &pd1, &pd2, &pd3, &pd4, &pd5)
}
