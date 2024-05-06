package repository

import (
	"database/sql"
	"fmt"

	"github.com/mrakhaf/enqilo-store/domain/auth-staff/interfaces"
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

func (r *repoHandler) SaveUserAccount(data request.Register) (id string, err error) {
	id = utils.GenerateUUID()

	hasPassword, _ := utils.HashPassword(data.Password)

	query := fmt.Sprintf(`INSERT INTO staff (id, phonenumber, password, name) VALUES ('%s', '%s', '%s', '%s')`, id, data.PhoneNumber, hasPassword, data.Name)

	_, err = r.databaseDB.Exec(query)

	if err != nil {
		return
	}

	return
}

func (r *repoHandler) GetDataAccount(phoneNumber string) (data entity.User, err error) {

	row := r.databaseDB.QueryRow(fmt.Sprintf("SELECT id, name, phoneNumber, password FROM staff WHERE phoneNumber = '%s'", phoneNumber))

	err = row.Scan(&data.Id, &data.Name, &data.PhoneNumber, &data.Password)

	if err != nil {
		return
	}

	return
}
