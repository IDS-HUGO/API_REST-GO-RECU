package application

//holaaaa
import (
	"demo/src/products/application/messaging"
	"demo/src/products/domain/entities"
	"demo/src/products/infraestructure/repositories"
)

type CreateProduct struct {
	db      repositories.ProductRepository
	publish messaging.PublishProductCreated
}

func NewCreateProduct(db repositories.ProductRepository, publish messaging.PublishProductCreated) CreateProduct {
	return CreateProduct{db: db, publish: publish}
}

func (cp *CreateProduct) Execute(product *entities.Product) error {
	err := cp.db.Save(product)
	if err != nil {
		return err
	}

	return cp.publish.Execute(product)
}
