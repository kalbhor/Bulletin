package Models

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name     string    `json:"Name"`
	ImageUrl string    `json:"ImageUrl"`
	Services []Service `json:"Services" gorm:"ForeignKey:CategoryRefer"`
}

type Service struct {
	gorm.Model
	Name     string  `json:"Name"`
	ImageUrl string  `json:"ImageUrl"`
	Offers   []Offer `json:"Offers" gorm:"ForeignKey:ServiceRefer"`
	CategoryRefer uint 
}

type Offer struct {
	gorm.Model
	Description string   `json:"Description"`
	Media       []Media `json:"Media"`
	Expectation      uint `json:"Expectation"`
	ServiceRefer uint 
	UserRefer uint 
	//Expectation

}

type OfferRequest struct {
	gorm.Model
	Description string   `json:"Description"`
	Media       []Media `json:"Media"`
	Budget      int      `json:"Budget"`
	UserRefer uint 
	//Expectation

}

type Media struct {
	gorm.Model
	ImageURL string `json:"ImageURL"`
	VideoURL string  `json:"VideoURL"`
	URL string `json:"URL"`
}

type User struct {
	gorm.Model
	Name string `json:"Name"`
	Email string `json:"Email"`
	Offers []Offer `gorm:"ForeignKey:UserRefer"`
	OfferRequests []OfferRequest `gorm:"ForeignKey:UserRefer"`
}

func (c *Category) TableName() string {
	return "category"
}

func (t *Service) TableName() string {
	return "service"
}

func (o *Offer) TableName() string {
	return "offer"
}