package Model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name     string `gorm:"not null;column:name"`
	Type     string `gorm:"not null;column:type"`
	Category string `gorm:"not null;column:category"`
	Price    int    `gorm:"not null;column:price"`
}

func NewProduct(Name, Type, Category string, Price int) *Product {
	p := &Product{
		Name:     Name,
		Type:     Type,
		Category: Category,
		Price:    Price,
	}
	return p
}

func (p *Product) GetID() uint {
	return p.ID
}
