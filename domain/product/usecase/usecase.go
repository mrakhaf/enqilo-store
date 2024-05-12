package usecase

import (
	"context"
	"fmt"

	customer "github.com/mrakhaf/enqilo-store/domain/customer/interfaces"
	"github.com/mrakhaf/enqilo-store/domain/product/interfaces"
	product "github.com/mrakhaf/enqilo-store/domain/product/interfaces"

	"github.com/mrakhaf/enqilo-store/models/entity"
	"github.com/mrakhaf/enqilo-store/models/request"
	"github.com/mrakhaf/enqilo-store/models/response"
)

type usecase struct {
	customerRepo customer.Repository
	productRepo  product.Repository
}

func NewUsecase(customerRepo customer.Repository, productRepo product.Repository) interfaces.Usecase {
	return &usecase{
		customerRepo: customerRepo,
		productRepo:  productRepo,
	}
}

func (u *usecase) CreateProduct(ctx context.Context, data request.CreateProduct) (id string, createdAt string, err error) {

	id, createdAt, err = u.productRepo.SaveProduct(data)

	if err != nil {
		err = fmt.Errorf("failed to save product: %s", err)
		return
	}

	return

}

func (u *usecase) Checkout(ctx context.Context, req request.Checkout) (id string, createdAt string, err error) {

	customer, _ := u.customerRepo.SearchCustomerAccount(req.CustomerId)

	if customer.Id == "" {
		err = fmt.Errorf("customer account not found")
		return
	}

	var total int

	for _, item := range req.ProductDetails {
		product, _ := u.productRepo.GetDataProductById(item.ProductId)
		price := product.Price * item.Quantity

		total += price

		if product.Id == "" {
			err = fmt.Errorf("product not found")
			return
		}

		if product.Stock < item.Quantity {
			err = fmt.Errorf("stock is not enough")
			return
		}

		if !product.IsAvailable {
			err = fmt.Errorf("product is not available")
			return
		}
	}

	if req.Paid < total {
		err = fmt.Errorf("paid is less than total")
		return
	}

	if total-req.Paid != req.Change {
		err = fmt.Errorf("change is not right")
		return
	}

	id, createdAt, err = u.productRepo.Checkout(req)

	if err != nil {
		err = fmt.Errorf("failed to checkout product: %s", err)
		return
	}

	return
}

func (u *usecase) SearchProducts(ctx context.Context, req request.GetProducts) (data []entity.Product, err error) {
	var firstFilterParam bool

	query := "SELECT * FROM products"

	if req.Id != nil {
		if !firstFilterParam {
			query = fmt.Sprintf("%s WHERE id = '%s' ", query, *req.Id)
			firstFilterParam = true
		} else {
			query = fmt.Sprintf("%s AND id = '%s' ", query, *req.Id)
		}
	}

	if req.Name != nil {
		if !firstFilterParam {
			query = fmt.Sprintf("%s WHERE LOWER(name) like '%%%s%%' ", query, *req.Name)
			firstFilterParam = true
		} else {
			query = fmt.Sprintf("%s AND WHERE LOWER(name) like '%%%s%%' ", query, *req.Name)
		}
	}

	if req.IsAvailable != nil && (*req.IsAvailable == "true" || *req.IsAvailable == "false") {
		if !firstFilterParam {
			if *req.IsAvailable == "true" {
				query = fmt.Sprintf("%s WHERE isavailable IS FALSE ", query)
			} else {
				query = fmt.Sprintf("%s WHERE isavailable IS TRUE", query)
			}
			firstFilterParam = true
		} else {
			if *req.IsAvailable == "true" {
				query = fmt.Sprintf("%s AND isavailable IS FALSE", query)
			} else {
				query = fmt.Sprintf("%s AND isavailable IS TRUE", query)
			}
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

	if req.InStock != nil && (*req.InStock == "true" || *req.InStock == "false") {
		if !firstFilterParam {
			if *req.InStock == "true" {
				query = fmt.Sprintf("%s WHERE stock > 0", query)
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

	if req.Price != nil && req.CreatedAt == nil {
		if *req.Price == "asc" {
			query = fmt.Sprintf("%s ORDER BY price ASC", query)
		} else if *req.Price == "desc" {
			query = fmt.Sprintf("%s ORDER BY price DESC", query)
		}
	}

	if req.CreatedAt != nil && req.Price == nil {
		if *req.CreatedAt == "asc" {
			query = fmt.Sprintf("%s ORDER BY createdat ASC", query)
		} else if *req.CreatedAt == "desc" {
			query = fmt.Sprintf("%s ORDER BY createdat DESC", query)
		}
	}

	if req.Price == nil && req.CreatedAt == nil {
		query = fmt.Sprintf("%s ORDER BY createdat DESC", query)
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

	fmt.Println(query, "ini query")
	data, err = u.productRepo.SearchProducts(query)

	if err != nil {
		err = fmt.Errorf("failed to search products: %s", err)
		return
	}

	productResponse := []entity.Product{}

	for _, product := range data {
		productResponse = append(productResponse, entity.Product{
			Id:        product.Id,
			Name:      product.Name,
			Sku:       product.Sku,
			Category:  product.Category,
			ImageUrl:  product.ImageUrl,
			Price:     product.Price,
			Stock:     product.Stock,
			Location:  product.Location,
			CreatedAt: product.CreatedAt,
		})
	}

	data = productResponse

	return
}

func (u *usecase) UpdateProduct(ctx context.Context, id string, req request.CreateProduct) (err error) {

	err = u.productRepo.UpdateProduct(id, req)

	if err != nil {
		err = fmt.Errorf("failed to update product: %s", err)
		return
	}

	return

}

func (u *usecase) DeleteProduct(ctx context.Context, id string) (err error) {
	err = u.productRepo.DeleteProduct(id)

	if err != nil {
		err = fmt.Errorf("failed to delete product: %s", err)
		return
	}

	return
}

func (u *usecase) SearchSku(req request.SearchProductParam) (data interface{}, err error) {

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
				// firstFilterParam = true
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

	products, err := u.productRepo.SearchSku(query)

	if err != nil {
		err = fmt.Errorf("failed to search product: %s", err)
		return
	}

	productsResponse := []response.Products{}

	for _, product := range products {
		productsResponse = append(productsResponse, response.Products{
			Id:        product.Id,
			Name:      product.Name,
			Sku:       product.Sku,
			Category:  product.Category,
			ImageUrl:  product.ImageUrl,
			Price:     product.Price,
			Stock:     product.Stock,
			Location:  product.Location,
			CreatedAt: product.CreatedAt.Format("2006-01-02"),
		})
	}

	data = productsResponse

	return
}

func (u *usecase) GetCheckoutHistory(ctx context.Context, req request.GetCheckoutHistoryParam) (data interface{}, err error) {

	query := "select c.id as \"transactionId\", c2.id as \"customerId\", p.id as \"productId\", ci.quantity, c.paid, c.\"change\", c.createdat " +
		"from checkout c " +
		"join customer c2 on c.customerid = c2.id " +
		"join checkout_item ci on c.id = ci.checkoutid " +
		"join products p on ci.productid = p.id "

	var firstFilterParam bool

	if req.CustomerId != nil {
		if !firstFilterParam {
			query = fmt.Sprintf("%s WHERE c2.id = '%s'", query, *req.CustomerId)
		} else {
			query = fmt.Sprintf("%s AND c2.id = '%s'", query, *req.CustomerId)
		}
	}

	if req.CreatedAt != nil {
		if *req.CreatedAt == "asc" {
			query = fmt.Sprintf("%s ORDER BY c.createdat ASC", query)
		} else if *req.CreatedAt == "desc" {
			query = fmt.Sprintf("%s ORDER BY c.createdat DESC", query)
		}
	} else {
		query = fmt.Sprintf("%s ORDER BY c.createdat DESC", query)
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

	checkoutHistories, err := u.productRepo.GetCheckoutHistory(query)

	if err != nil {
		err = fmt.Errorf("failed to get history checkout: %s", err)
		return
	}

	checkoutHistoriesResponse := []response.CheckoutHistories{}

	for _, item := range checkoutHistories {
		checkoutHistoriesResponse = append(checkoutHistoriesResponse, response.CheckoutHistories{
			TransactionId: item.TransactionId,
			CustomerId:    item.CustomerId,
			ProductDetails: []response.ProductDetails{
				{
					ProductId: item.ProductId,
					Quantity:  item.Quantity,
				},
			},
			Quantity:  item.Quantity,
			Paid:      item.Paid,
			Change:    item.Change,
			CreatedAt: item.CreatedAt,
		})
	}

	data = checkoutHistoriesResponse

	return
}
