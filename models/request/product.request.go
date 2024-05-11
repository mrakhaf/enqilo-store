package request

type CreateProduct struct {
	Name        string `json:"name" validate:"required,min=1,max=30"`
	Sku         string `json:"sku" validate:"required,min=1,max=30"`
	Category    string `json:"category" validate:"required,oneof=Clothing Accessories Footwear Beverages"`
	ImageUrl    string `json:"imageUrl" validate:"required,url"`
	Notes       string `json:"notes" validate:"required,min=1,max=200"`
	Price       int    `json:"price" validate:"required,min=1"`
	Stock       *int   `json:"stock" validate:"required,gte=0,lte=100000"`
	Location    string `json:"location" validate:"required,min=1,max=200"`
	IsAvailable *bool  `json:"isAvailable" validate:"required"`
}

type ProductDetails struct {
	ProductId string `json:"productId" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required,gte=0,lte=100000"`
}

type Checkout struct {
	CustomerId     string           `json:"customerId" validate:"required"`
	ProductDetails []ProductDetails `json:"productDetails" validate:"required"`
	Paid           int              `json:"paid" validate:"required"`
	Change         int              `json:"change" validate:"min=0"`
}

type GetProducts struct {
	Id          *string `query:"id" validate:"omitempty"`
	Limit       *int    `query:"limit" validate:"omitempty,gte=0,lte=5"`
	Offset      *int    `query:"offset" validate:"omitempty,gte=0"`
	Name        *string `query:"name" validate:"omitempty"`
	IsAvailable *string `query:"isAvailable" validate:"omitempty"`
	Category    *string `query:"category" validate:"omitempty"`
	Sku         *string `query:"sku" validate:"omitempty"`
	Price       *string `query:"price" validate:"omitempty"`
	InStock     *string `query:"inStock" validate:"omitempty"`
	CreatedAt   *string `query:"createdAt" validate:"omitempty"`
}

type SearchProductParam struct {
	Limit    *int    `query:"limit" validate:"omitempty,gte=0,lte=5"`
	Offset   *int    `query:"offset" validate:"omitempty,gte=0"`
	Name     *string `query:"name" validate:"omitempty"`
	Category *string `query:"category" validate:"omitempty"`
	Sku      *string `query:"sku" validate:"omitempty"`
	Price    *string `query:"price" validate:"omitempty"`
	InStock  *string `query:"inStock" validate:"omitempty"`
}

type GetCheckoutHistoryParam struct {
	CustomerId *string `query:"customerId" validate:"omitempty"`
	Limit      *int    `query:"limit" validate:"omitempty,gte=0,lte=5"`
	Offset     *int    `query:"offset" validate:"omitempty,gte=0"`
	CreatedAt  *string `query:"createdAt" validate:"omitempty"`
}
