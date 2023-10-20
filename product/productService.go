package main

type ServiceInterface interface {
	Create(*Product) error
	Update(int, Product) error
	Delete(int) error
	Find(int) (*Product, error)
	List() ([]*Product, error)
}

type ProductService struct{}

func NewProductService() ProductService {
	return ProductService{}
}

func (ProductService) Create(*Product) error {
	//insert to DB(product)
	return nil
}
func (ProductService) Update(int, Product) error {
	//update(id)
	return nil
}
func (ProductService) Delete(int) error {
	//delete(id)
	return nil
}
func (ProductService) Find(int) (*Product, error) {
	//find(id)
	return nil, nil
}
func (ProductService) List() ([]*Product, error) {
	//list(id)
	return nil, nil
}
