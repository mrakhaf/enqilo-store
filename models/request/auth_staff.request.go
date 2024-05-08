package request

type Login struct {
	PhoneNumber string `json:"phoneNumber" validate:"required,min=10,max=16,startswith=+"`
	Password    string `json:"password" validate:"required,min=5,max=15"`
}

type Register struct {
	PhoneNumber string `json:"phoneNumber" validate:"required,min=10,max=16,startswith=+"`
	Password    string `json:"password" validate:"required,min=5,max=15"`
	Name        string `json:"name" validate:"required,min=5,max=50"`
}
