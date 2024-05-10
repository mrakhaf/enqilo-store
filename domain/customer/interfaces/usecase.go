package interfaces

import (
	"context"

	"github.com/mrakhaf/enqilo-store/models/request"
	"github.com/mrakhaf/enqilo-store/models/response"
)

type Usecase interface {
	Register(ctx context.Context, req request.RegisterCustomer) (data response.CustomerResponse, err error)
	GetCustomers(ctx context.Context, query request.GetAllCustomerParam) (data interface{}, err error)
}
