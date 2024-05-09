package usecase

import (
	"context"
	"fmt"

	"github.com/mrakhaf/enqilo-store/domain/customer/interfaces"
	"github.com/mrakhaf/enqilo-store/models/request"
	"github.com/mrakhaf/enqilo-store/models/response"
)

type (
	usecase struct {
		repository interfaces.Repository
	}
)

func NewUsecase(repository interfaces.Repository) interfaces.Usecase {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) Register(ctx context.Context, req request.RegisterCustomer) (data response.CustomerResponse, err error) {
	isExist, _ := u.repository.GetDataAccount(req.PhoneNumber)

	if isExist.PhoneNumber != "" {
		err = fmt.Errorf("phone number already exist")
		return
	}

	//save db
	idCustomer, err := u.repository.SaveCustomerAccount(req)

	if err != nil {
		err = fmt.Errorf("failed to save custoner account: %s", err)
		return
	}

	data = response.CustomerResponse{
		PhoneNumber: req.PhoneNumber,
		Id:          idCustomer,
		Name:        req.Name,
	}

	return
}
