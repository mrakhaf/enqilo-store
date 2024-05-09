package response

type Products struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Sku       string `json:"sku"`
	Category  string `json:"category"`
	ImageUrl  string `json:"imageUrl"`
	Price     int    `json:"price"`
	Stock     int    `json:"stock"`
	Location  string `json:"location"`
	CreatedAt string `json:"createdAt"`
}
