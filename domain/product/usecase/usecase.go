package usecase

import (
	"context"
	"fmt"

	"github.com/mrakhaf/enqilo-store/domain/product/interfaces"
	"github.com/mrakhaf/enqilo-store/models/entity"
	"github.com/mrakhaf/enqilo-store/models/request"
	"github.com/mrakhaf/enqilo-store/shared/utils"
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

func (u *usecase) SearchProducts(ctx context.Context, req request.GetProducts) (data []entity.Product, err error) {

	query := "SELECT * FROM products WHERE 1=1 "

	if req.Id != nil {
		query = fmt.Sprintf("%s AND id = '%s' ", query, *req.Id)
	}

	if req.Name != nil {
		query = fmt.Sprintf("%s AND name like '%%%s%%' ", query, *req.Name)
	}

	if req.IsAvailable != nil {
		if *req.IsAvailable == "true" {
			query = fmt.Sprintf("%s AND isAvailable IS TRUE ", query)
		} else if *req.IsAvailable == "false" {
			query = fmt.Sprintf("%s AND isAvailable IS FALSE ", query)
		}
	}

	if req.Category != nil {
		categories := []string{"Clothing", "Accessories", "Footwear", "Beverages"}

		if utils.Contains(categories, *req.Category) {
			query = fmt.Sprintf("%s AND category = '%s' ", query, *req.Category)
		}
	}

	if req.Sku != nil {
		query = fmt.Sprintf("%s AND sku = '%s' ", query, *req.Sku)
	}

	if req.InStock != nil {
		if *req.InStock == "true" {
			query = fmt.Sprintf("%s AND stock > 0 ", query)
		} else if *req.InStock == "false" {
			query = fmt.Sprintf("%s AND stock = 0 ", query)
		}
	}

	if req.Price != nil {
		if *req.Price == "asc" {
			query = fmt.Sprintf("%s ORDER BY price ASC ", query)
		} else if *req.Price == "desc" {
			query = fmt.Sprintf("%s ORDER BY price DESC ", query)
		}
	}

	if req.CreatedAt != nil {
		if *req.CreatedAt == "asc" {
			query = fmt.Sprintf("%s ORDER BY createdat ASC ", query)
		} else if *req.CreatedAt == "desc" {
			query = fmt.Sprintf("%s ORDER BY createdat DESC ", query)
		}
	}

	fmt.Println(query)

	data, err = u.repository.SearchProducts(query)

	if err != nil {
		err = fmt.Errorf("failed to search products: %s", err)
		return
	}

	return
}

func (u *usecase) UpdateProduct(ctx context.Context, id string, req request.CreateProduct) (err error) {

	err = u.repository.UpdateProduct(id, req)

	if err != nil {
		err = fmt.Errorf("failed to update product: %s", err)
		return
	}

	return

}

func (u *usecase) DeleteProduct(ctx context.Context, id string) (err error) {
	err = u.repository.DeleteProduct(id)

	if err != nil {
		err = fmt.Errorf("failed to delete product: %s", err)
		return
	}

	return
}
