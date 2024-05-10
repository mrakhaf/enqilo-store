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

type ProductDetails struct {
	ProductId string `json:"productId" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required,gte=0,lte=100000"`
}

type CheckoutHistory struct {
	TransactionId string `json:"transactionId"`
	CustomerId    string `json:"customerId"`
	ProductId     string `json:"productId" validate:"required"`
	Quantity      int    `json:"quantity" validate:"required,gte=0,lte=100000"`
	Paid          int    `json:"paid"`
	Change        int    `json:"change"`
	CreatedAt     string `json:"createdAt"`
}

type CheckoutHistories struct {
	TransactionId  string           `json:"transactionId"`
	CustomerId     string           `json:"customerId"`
	ProductDetails []ProductDetails `json:"productDetails"`
	Quantity       int              `json:"quantity" validate:"required,gte=0,lte=100000"`
	Paid           int              `json:"paid"`
	Change         int              `json:"change"`
	CreatedAt      string           `json:"createdAt"`
}
