package interfaces

import "github.com/mrakhaf/enqilo-store/models/request"

type Repository interface {
	SaveProduct(data request.CreateProduct) (id string, createdAt string, err error)
}
