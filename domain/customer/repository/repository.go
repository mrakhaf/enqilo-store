package repository

import (
	"database/sql"
	"fmt"

	"github.com/mrakhaf/enqilo-store/domain/customer/interfaces"
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

func (r *repoHandler) SaveCustomerAccount(data request.RegisterCustomer) (id string, err error) {
	id = utils.GenerateUUID()

	query := fmt.Sprintf(`INSERT INTO customer (id, phonenumber, name) VALUES ('%s', '%s', '%s')`, id, data.PhoneNumber, data.Name)

	_, err = r.databaseDB.Exec(query)

	if err != nil {
		return
	}

	return
}

func (r *repoHandler) GetDataAccount(phoneNumber string) (data entity.Customer, err error) {
	row := r.databaseDB.QueryRow(fmt.Sprintf("SELECT phoneNumber FROM customer WHERE phoneNumber = '%s'", phoneNumber))

	err = row.Scan(&data.PhoneNumber)

	if err != nil {
		return
	}

	return
}

func (r *repoHandler) SearchCustomerAccount(id string) (data entity.Customer, err error) {
	row := r.databaseDB.QueryRow(fmt.Sprintf("SELECT id, phonenumber, name FROM customer WHERE id = '%s'", id))

	err = row.Scan(&data.Id, &data.PhoneNumber, &data.Name)

	if err != nil {
		return
	}

	return
}

func (r *repoHandler) GetAllCustomer(query string) (customers []entity.Customer, err error) {

	rows, err := r.databaseDB.Query(query)

	if err != nil {
		return
	}

	defer rows.Close()

	customer := entity.Customer{}

	for rows.Next() {

		err = rows.Scan(&customer.Id, &customer.PhoneNumber, &customer.Name)

		customers = append(customers, customer)
	}

	return
}
