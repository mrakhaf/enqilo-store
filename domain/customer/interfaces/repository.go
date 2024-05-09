package interfaces

import (
	"github.com/mrakhaf/enqilo-store/models/entity"
	"github.com/mrakhaf/enqilo-store/models/request"
)

type Repository interface {
	SaveCustomerAccount(data request.RegisterCustomer) (id string, err error)
	GetDataAccount(phoneNumber string) (data entity.Customer, err error)
	SearchCustomerAccount(id string) (data entity.Customer, err error)
}
