package entity

import "time"

type User struct {
	Id          string
	PhoneNumber string
	Password    string
	Name        string
}

type Customer struct {
	Id          string
	PhoneNumber string
	Name        string
}

type Product struct {
	Id          string    `json:"id"`
	Name        string    `json:"name" validate:"required,min=1,max=30"`
	Sku         string    `json:"sku" validate:"required,min=1,max=30"`
	Category    string    `json:"category" validate:"required,oneof=Clothing Accessories Footwear Beverages"`
	ImageUrl    string    `json:"imageUrl" validate:"required,url"`
	Notes       string    `json:"notes" validate:"required,min=1,max=200"`
	Price       int       `json:"price" validate:"required"`
	Stock       int       `json:"stock" validate:"required,gte=0,lte=100000"`
	Location    string    `json:"location" validate:"required,min=1,max=200"`
	IsAvailable bool      `json:"isAvailable" validate:"required"`
	CreatedAt   time.Time `json:"createdAt"`
}
