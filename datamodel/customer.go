package datamodel

import (
	"time"
)

// type stringstring
type Customer struct {
	// Uuid      string     `json:"uuid" gorm:"primaryKey;column:uuid;not null"`
	Tel        string    `json:"tel" gorm:"primaryKey;column:tel;not null"`
	Name       string    `json:"name" gorm:"column:name;not null"`
	Email      string    `json:"email" gorm:"column:email;not null"`
	Social     string    `json:"social_network" gorm:"column:social_network"`
	SocialData string    `json:"-"`
	CreatedAt  time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"-" gorm:"autoUpdateTime"`
}

type SocialData map[string]string

func NewCustomer(name string, tel string, email string, social string) *Customer {
	return &Customer{
		// Uuid:      string(uuid.NewString()),
		Tel:    tel,
		Name:   name,
		Email:  email,
		Social: social,
	}
}

func (Customer) TableName() string {
	return "film_custs"
}
