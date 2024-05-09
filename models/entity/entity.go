package entity

import "time"

type User struct {
	Id          string
	PhoneNumber string
	Password    string
	Name        string
}

type Product struct {
	Id          string
	Name        string
	Sku         string
	Category    string
	ImageUrl    string
	Notes       string
	Price       int
	Stock       int
	Location    string
	IsAvailable bool
	CreatedAt   time.Time
}
