package interfaces

import (
	"context"

	"github.com/mrakhaf/enqilo-store/models/request"
)

type Usecase interface {
	CreateProduct(ctx context.Context, data request.CreateProduct) (id string, createdAt string, err error)
}
