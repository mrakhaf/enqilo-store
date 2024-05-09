package repository

import (
	"database/sql"
	"fmt"
	"log"
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

func (r *repoHandler) SearchProducts(query string) (data []entity.Product, err error) {

	row, err := r.databaseDB.Query(query)

	if err != nil {
		if err == sql.ErrNoRows {
			return data, nil
		}
		return
	}

	for row.Next() {
		var id, name, sku, category, imageurl, notes, location string
		var price, stock int
		var isAvailable bool
		var createdAt string

		err = row.Scan(&id, &name, &sku, &category, &imageurl, &notes, &price, &stock, &location, &isAvailable, &createdAt)

		if err != nil {
			return
		}

		data = append(data, entity.Product{
			Id:          id,
			Name:        name,
			Sku:         sku,
			Category:    category,
			ImageUrl:    imageurl,
			Notes:       notes,
			Price:       price,
			Stock:       stock,
			Location:    location,
			IsAvailable: isAvailable,
			CreatedAt:   createdAt,
		})
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

	_, err = r.databaseDB.Exec(fmt.Sprintf("UPDATE products SET name = '%s', sku = '%s', category = '%s', imageurl = '%s', notes = '%s', price = '%d', stock = '%d', location = '%s', isAvailable = '%t' WHERE id = '%s'", req.Name, req.Sku, req.Category, req.ImageUrl, req.Notes, req.Price, req.Stock, req.Location, req.IsAvailable, id))

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
