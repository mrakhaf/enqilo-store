package interfaces

import (
	"context"

	"github.com/mrakhaf/enqilo-store/models/entity"
	"github.com/mrakhaf/enqilo-store/models/request"
)

type Usecase interface {
	CreateProduct(ctx context.Context, data request.CreateProduct) (id string, createdAt string, err error)
	SearchProducts(ctx context.Context, req request.GetProducts) (data []entity.Product, err error)
	UpdateProduct(ctx context.Context, id string, req request.CreateProduct) (err error)
	DeleteProduct(ctx context.Context, id string) (err error)
}
