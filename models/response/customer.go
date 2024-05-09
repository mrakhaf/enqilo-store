package response

type CustomerResponse struct {
	PhoneNumber string `json:"phoneNumber"`
	Id          string `json:"userId"`
	Name        string `json:"name"`
}
