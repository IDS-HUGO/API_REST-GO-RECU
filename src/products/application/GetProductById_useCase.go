package application

import (
	"demo/src/products/domain/entities"
	"demo/src/products/infraestructure/repositories"
)

type GetProductById struct {
	db repositories.ProductRepository
}

func NewGetProductById(db repositories.ProductRepository) *GetProductById {
	return &GetProductById{db: db}
}

func (gp *GetProductById) Execute(id int) (*entities.Product, error) {
	product, err := gp.db.GetProductById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
