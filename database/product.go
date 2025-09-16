package database

var productList []Product

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func Store(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

func List() []Product {
	return productList
}

func Get(productId int) *Product {
	for _, p := range productList {
		if p.ID == productId {
			return &p
		}
	}
	return nil
}

func Update(product Product) {
	for idx, p := range productList {
		if p.ID == product.ID {
			productList[idx] = product
		}
	}
}

func Delete(productId int) {
	for idx, p := range productList {
		if p.ID == productId {
			productList = append(productList[:idx], productList[idx+1:]...)
		}
	}
}

func init() {
	pd1 := Product{ID: 1, Name: "Laptop", Price: 1000}
	pd2 := Product{ID: 2, Name: "Laptop", Price: 1000}
	pd3 := Product{ID: 3, Name: "Laptop", Price: 1000}
	pd4 := Product{ID: 4, Name: "Laptop", Price: 1000}
	pd5 := Product{ID: 5, Name: "Laptop", Price: 1000}

	productList = append(productList, pd1, pd2, pd3, pd4, pd5)
}
