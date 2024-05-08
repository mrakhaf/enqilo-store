package usecase

import (
	"context"
	"fmt"

	"github.com/mrakhaf/enqilo-store/domain/product/interfaces"
	"github.com/mrakhaf/enqilo-store/models/request"
)

type usecase struct {
	repository interfaces.Repository
}

func NewUsecase(repository interfaces.Repository) interfaces.Usecase {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) CreateProduct(ctx context.Context, data request.CreateProduct) (id string, createdAt string, err error) {

	id, createdAt, err = u.repository.SaveProduct(data)

	if err != nil {
		err = fmt.Errorf("failed to save product: %s", err)
		return
	}

	return

}

func (u *usecase) ValidateQueryParams(ctx context.Context, data request.GetProducts) (query string, err error) {

	if data.Sku != nil {

	}

	return
}
