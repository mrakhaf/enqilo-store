package request

type CreateProduct struct {
	Name        string `json:"name" validate:"required,min=1,max=30"`
	Sku         string `json:"sku" validate:"required,min=1,max=30"`
	Category    string `json:"category" validate:"required,oneof=Clothing Accessories Footwear Beverages"`
	ImageUrl    string `json:"imageUrl" validate:"required,url"`
	Notes       string `json:"notes" validate:"required,min=1,max=200"`
	Price       int    `json:"price" validate:"required"`
	Stock       int    `json:"stock" validate:"required,gte=0,lte=100000"`
	Location    string `json:"location" validate:"required,min=1,max=200"`
	IsAvailable bool   `json:"isAvailable" validate:"required"`
}

type GetProducts struct {
	Id          *string `json:"id" validate:"omitempty"`
	Limit       *int    `json:"limit" validate:"omitempty,gte=0,lte=5"`
	Offset      *int    `json:"offset" validate:"omitempty,gte=0"`
	Name        *string `json:"name" validate:"omitempty"`
	IsAvailable *string `json:"isAvailable" validate:"omitempty"`
	Category    *string `json:"category" validate:"omitempty"`
	Sku         *string `json:"sku" validate:"omitempty"`
	Price       *string `json:"price" validate:"omitempty"`
	InStock     *string `json:"inStock" validate:"omitempty"`
	CreatedAt   *string `json:"createdAt" validate:"omitempty"`
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
