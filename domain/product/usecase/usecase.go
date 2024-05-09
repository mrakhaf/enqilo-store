package usecase

import (
	"context"
	"fmt"

	"github.com/mrakhaf/enqilo-store/domain/product/interfaces"
	"github.com/mrakhaf/enqilo-store/models/request"
	"github.com/mrakhaf/enqilo-store/models/response"
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

func (u *usecase) SearchProduct(req request.SearchProductParam) (data interface{}, err error) {

	query := "SELECT id, name, sku, category, imageurl, price, stock, location, createdat FROM products"

	var firstFilterParam bool

	if req.Name != nil {
		if !firstFilterParam {
			query = fmt.Sprintf("%s WHERE LOWER(name) LIKE '%%%s%%'", query, *req.Name)
			firstFilterParam = true
		} else {
			query = fmt.Sprintf("%s AND LOWER(name) LIKE '%%%s%%'", query, *req.Name)
		}
	}

	if req.Category != nil {
		category := *req.Category
		switch category {
		case "Clothing", "Accessories", "Footwear", "Beverages":
			if !firstFilterParam {
				query = fmt.Sprintf("%s WHERE category = '%s'", query, category)
				firstFilterParam = true
			} else {
				query = fmt.Sprintf("%s AND category = '%s'", query, category)
			}
		}
	}

	if req.Sku != nil {
		if !firstFilterParam {
			query = fmt.Sprintf("%s WHERE LOWER(sku) LIKE '%%%s%%'", query, *req.Sku)
			firstFilterParam = true
		} else {
			query = fmt.Sprintf("%s AND LOWER(sku) LIKE '%%%s%%'", query, *req.Sku)
		}
	}

	if req.Price != nil {
		if *req.Price == "asc" {
			query = fmt.Sprintf("%s ORDER BY price ASC", query)
		} else if *req.Price == "desc" {
			query = fmt.Sprintf("%s ORDER BY price DESC", query)
		}
	}

	if req.InStock != nil && (*req.InStock == "true" || *req.InStock == "false") {
		if !firstFilterParam {
			if *req.InStock == "true" {
				query = fmt.Sprintf("%s WHERE stock > 0", query)
				firstFilterParam = true
			} else {
				query = fmt.Sprintf("%s WHERE stock <= 0", query)
			}
		} else {
			if *req.InStock == "true" {
				query = fmt.Sprintf("%s AND stock > 0", query)
			} else {
				query = fmt.Sprintf("%s AND stock <= 0", query)
			}
		}
	}

	if req.Limit != nil {
		query = fmt.Sprintf("%s LIMIT %d", query, *req.Limit)
	} else {
		query = fmt.Sprintf("%s LIMIT 5", query)
	}

	if req.Offset != nil {
		query = fmt.Sprintf("%s OFFSET %d", query, *req.Offset)
	} else {
		query = fmt.Sprintf("%s OFFSET 0", query)
	}

	products, err := u.repository.SearchProduct(query)

	if err != nil {
		err = fmt.Errorf("failed to search product: %s", err)
		return
	}

	productsResponse := []response.Products{}

	for _, product := range products {
		createdAt := product.CreatedAt.Format("2006-01-02")

		productsResponse = append(productsResponse, response.Products{
			Id:        product.Id,
			Name:      product.Name,
			Sku:       product.Sku,
			Category:  product.Category,
			ImageUrl:  product.ImageUrl,
			Price:     product.Price,
			Stock:     product.Stock,
			Location:  product.Location,
			CreatedAt: createdAt,
		})
	}

	data = productsResponse

	return
}
