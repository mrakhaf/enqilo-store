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
