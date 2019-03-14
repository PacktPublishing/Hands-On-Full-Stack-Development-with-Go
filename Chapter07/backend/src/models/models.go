package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Image       string  `json:"img"`
	ImagAlt     string  `json:"imgalt" gorm:"column:imgalt"`
	Price       float64 `json:"price"`
	Promotion   float64 `json:"promotion"`
	PoructName  string  `gorm:"column:productname" json:"productname"`
	Description string
}

func (Product) TableName() string {
	return "products"
}

type Customer struct {
	gorm.Model
	FirstName string `gorm:"column:firstname" json:"firstname"`
	LastName  string `gorm:"column:lastname" json:"lastname"`
	Email     string `gorm:"column:email" json:"email"`
	Pass      string `json:"password"`
	CCToken   string `gorm:"column:cctoken" json:"cctoken"`
	LoggedIn  bool   `gorm:"column:loggedin" json:"loggedin"`
}

func (Customer) TableName() string {
	return "customers"
}

type Order struct {
	gorm.Model
	Product
	Customer
	CustomerID   int       `gorm:"column:customer_id"`
	ProductID    int       `gorm:"column:product_id"`
	Price        float64   `gorm:"column:price" json:"sell_price"`
	PurchaseDate time.Time `gorm:"column:purchase_date" json:"purchase_date"`
}

func (Order) TableName() string {
	return "orders"
}
