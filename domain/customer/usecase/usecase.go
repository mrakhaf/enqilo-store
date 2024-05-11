package usecase

import (
	"context"
	"fmt"

	"github.com/mrakhaf/enqilo-store/domain/customer/interfaces"
	"github.com/mrakhaf/enqilo-store/models/request"
	"github.com/mrakhaf/enqilo-store/models/response"
	"github.com/mrakhaf/enqilo-store/shared/utils"
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
	utils.IsValidPhoneNumber(req.PhoneNumber)

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

func (u *usecase) GetCustomers(ctx context.Context, req request.GetAllCustomerParam) (data interface{}, err error) {
	query := "SELECT id, phonenumber, name FROM customer"

	var firstFilterParam bool

	if req.Name != nil {
		if !firstFilterParam {
			query = fmt.Sprintf("%s WHERE LOWER(name) LIKE '%%%s%%'", query, *req.Name)
			firstFilterParam = true
		} else {
			query = fmt.Sprintf("%s AND LOWER(name) LIKE '%%%s%%'", query, *req.Name)
		}
	}

	if req.PhoneNumber != nil {
		if !firstFilterParam {
			query = fmt.Sprintf("%s WHERE phoneNumber LIKE '+%%%s%%'", query, *req.PhoneNumber)
			firstFilterParam = true
		} else {
			query = fmt.Sprintf("%s AND phoneNumber LIKE '+%%%s%%'", query, *req.PhoneNumber)
		}
	}

	customers, err := u.repository.GetAllCustomer(query)

	if err != nil {
		err = fmt.Errorf("failed to get customers: %s", err)
		return
	}

	customerResponses := []response.CustomerResponse{}

	for _, item := range customers {
		customerResponses = append(customerResponses, response.CustomerResponse{
			Id:          item.Id,
			PhoneNumber: item.PhoneNumber,
			Name:        item.Name,
		})
	}

	data = customerResponses

	return
}
