package interfaces

import (
	"github.com/mrakhaf/enqilo-store/models/entity"
	"github.com/mrakhaf/enqilo-store/models/request"
)

type Repository interface {
	SaveProduct(data request.CreateProduct) (id string, createdAt string, err error)
	SearchSku(query string) (products []entity.Product, err error)
	Checkout(data request.Checkout) (id string, createdAt string, err error)
	SearchProducts(query string) (data []entity.Product, err error)
	GetDataProductById(id string) (data entity.Product, err error)
	UpdateProduct(id string, req request.CreateProduct) (err error)
	DeleteProduct(id string) (err error)
}
