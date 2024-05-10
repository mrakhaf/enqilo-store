package request

type RegisterCustomer struct {
	PhoneNumber string `json:"phoneNumber" validate:"required,min=10,max=16,startswith=+"`
	Name        string `json:"name" validate:"required,min=5,max=50"`
}

type GetAllCustomerParam struct {
	Name        *string `query:"name" validate:"omitempty"`
	PhoneNumber *string `query:"phoneNumber" validate:"omitempty"`
}
