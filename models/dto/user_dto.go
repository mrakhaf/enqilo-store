package dto

type AuthResponse struct {
	PhoneNumber string `json:"phoneNumber"`
	AccessToken string `json:"accessToken"`
	Id          string `json:"id"`
	Name        string `json:"name"`
}
