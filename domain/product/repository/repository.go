package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/mrakhaf/enqilo-store/domain/product/interfaces"
	"github.com/mrakhaf/enqilo-store/models/entity"
	"github.com/mrakhaf/enqilo-store/models/request"
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

	query := fmt.Sprintf(`INSERT INTO products (id, name, sku, category, imageurl, notes, price, stock, location, isAvailable, createdAt) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', %d, %d, '%s', %t, '%s')`, id, data.Name, data.Sku, data.Category, data.ImageUrl, data.Notes, data.Price, data.Stock, data.Location, data.IsAvailable, timeNow.Format("2006-01-02 15:04:05"))

	_, err = r.databaseDB.Exec(query)

	if err != nil {
		return
	}

	createdAt = timeNow.Format("2006-01-02")

	return
}

func (r *repoHandler) SearchProduct(query string) (products []entity.Product, err error) {

	rows, err := r.databaseDB.Query(query)

	if err != nil {
		return
	}

	defer rows.Close()

	product := entity.Product{}

	for rows.Next() {

		err = rows.Scan(&product.Id, &product.Name, &product.Sku, &product.Category, &product.ImageUrl, &product.Price, &product.Stock, &product.Location, &product.CreatedAt)

		if err != nil {
			return
		}

		products = append(products, product)
	}

	return
}
