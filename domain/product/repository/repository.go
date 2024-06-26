package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/mrakhaf/enqilo-store/domain/product/interfaces"
	"github.com/mrakhaf/enqilo-store/models/entity"
	"github.com/mrakhaf/enqilo-store/models/request"
	"github.com/mrakhaf/enqilo-store/models/response"
	"github.com/mrakhaf/enqilo-store/shared/utils"
)

type repoHandler struct {
	databaseDB *sql.DB
}

func NewRepository(databaseDB *sql.DB) interfaces.Repository {
	return &repoHandler{
		databaseDB: databaseDB,
	}
}

func (r *repoHandler) SaveProduct(data request.CreateProduct) (id string, createdAt string, err error) {

	id = utils.GenerateUUID()
	timeNow := time.Now()

	query := fmt.Sprintf(`INSERT INTO products (id, name, sku, category, imageurl, notes, price, stock, location, isAvailable, createdAt) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', %d, %d, '%s', %t, '%s')`, id, data.Name, data.Sku, data.Category, data.ImageUrl, data.Notes, data.Price, data.Stock, data.Location, *data.IsAvailable, timeNow.Format("2006-01-02 15:04:05"))

	_, err = r.databaseDB.Exec(query)

	if err != nil {
		return
	}

	createdAt = timeNow.Format("2006-01-02")

	return
}

func (r *repoHandler) Checkout(data request.Checkout) (id string, createdAt string, err error) {

	idCheckout := utils.GenerateUUID()
	timeNow := time.Now()

	query := fmt.Sprintf(`INSERT INTO checkout (id, customerId, paid, change, createdAt) VALUES ('%s', '%s', %d, %d, '%s')`, idCheckout, data.CustomerId, data.Paid, data.Change, timeNow.Format("2006-01-02 15:04:05"))
	_, err = r.databaseDB.Exec(query)

	if err != nil {
		return
	}

	for _, product := range data.ProductDetails {
		idItem := utils.GenerateUUID()
		queryCheckoutItem := fmt.Sprintf(`INSERT INTO checkout_item (id, checkoutId, productId, quantity, createdAt) VALUES ('%s', '%s', '%s', '%d', '%s')`, idItem, idCheckout, product.ProductId, product.Quantity, timeNow.Format("2006-01-02 15:04:05"))
		_, err = r.databaseDB.Exec(queryCheckoutItem)
		if err != nil {
			return
		}
	}

	createdAt = timeNow.Format("2006-01-02")

	return idCheckout, createdAt, nil
}

func (r *repoHandler) SearchSku(query string) (products []entity.Product, err error) {

	rows, err := r.databaseDB.Query(query)

	if err != nil {
		return
	}

	defer rows.Close()

	product := entity.Product{}

	for rows.Next() {

		err = rows.Scan(&product.Id, &product.Name, &product.Sku, &product.Category, &product.ImageUrl, &product.Price, &product.Stock, &product.Location, &product.CreatedAt)

		products = append(products, product)
	}

	return
}

func (r *repoHandler) SearchProducts(query string) (data []entity.Product, err error) {
	rows, err := r.databaseDB.Query(query)

	if err != nil {
		return
	}

	defer rows.Close()

	product := entity.Product{}

	for rows.Next() {

		err = rows.Scan(&product.Id, &product.Name, &product.Sku, &product.Category, &product.ImageUrl, &product.Notes, &product.Price, &product.Stock, &product.Location, &product.IsAvailable, &product.CreatedAt)

		data = append(data, product)
	}

	return
}

func (r *repoHandler) GetDataProductById(id string) (data entity.Product, err error) {
	row := r.databaseDB.QueryRow(fmt.Sprintf("SELECT id, name, sku, category, imageurl, notes, price, stock, location, isAvailable, createdAt FROM products WHERE id = '%s'", id))

	err = row.Scan(&data.Id, &data.Name, &data.Sku, &data.Category, &data.ImageUrl, &data.Notes, &data.Price, &data.Stock, &data.Location, &data.IsAvailable, &data.CreatedAt)

	if err != nil {
		return
	}

	return
}

func (r *repoHandler) UpdateProduct(id string, req request.CreateProduct) (err error) {

	query := fmt.Sprintf("UPDATE products SET name = '%s', sku = '%s', category = '%s', imageurl = '%s', notes = '%s', price = '%d', stock = '%d', location = '%s', isAvailable = '%t' WHERE id = '%s'", req.Name, req.Sku, req.Category, req.ImageUrl, req.Notes, req.Price, req.Stock, req.Location, *req.IsAvailable, id)

	fmt.Println(query)
	_, err = r.databaseDB.Exec(query)

	if err != nil {
		return
	}

	return
}

func (r *repoHandler) DeleteProduct(id string) (err error) {

	_, err = r.databaseDB.Exec(fmt.Sprintf("DELETE FROM products WHERE id = '%s'", id))

	if err != nil {
		log.Printf("failed to delete product: %s", err)
		return
	}

	return
}

func (r *repoHandler) GetCheckoutHistory(query string) (products []response.CheckoutHistory, err error) {

	rows, err := r.databaseDB.Query(query)

	if err != nil {
		return
	}

	defer rows.Close()

	checkoutHistories := response.CheckoutHistory{}

	for rows.Next() {

		err = rows.Scan(&checkoutHistories.TransactionId, &checkoutHistories.CustomerId, &checkoutHistories.ProductId, &checkoutHistories.Quantity, &checkoutHistories.Paid, &checkoutHistories.Change, &checkoutHistories.CreatedAt)

		products = append(products, checkoutHistories)
	}

	return
}
