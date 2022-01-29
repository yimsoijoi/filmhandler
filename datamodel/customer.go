package datamodel

import (
	"time"
)

type Info string
type Customer struct {
	// Uuid      Info      `json:"uuid" gorm:"primaryKey;column:uuid;not null"`
	Tel        Info      `json:"tel" gorm:"primaryKey;column:tel;not null"`
	Name       Info      `json:"name" gorm:"column:name;not null"`
	Email      Info      `json:"email" gorm:"column:email;not null"`
	Social     Info      `json:"social_network" gorm:"column:social_network"`
	SocialData string    `json:"-"`
	CreatedAt  time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"-" gorm:"autoUpdateTime"`
}

type SocialData map[string]string

func NewCustomer(name Info, tel Info, email Info, social Info) *Customer {
	return &Customer{
		// Uuid:      Info(uuid.NewString()),
		Tel:    tel,
		Name:   name,
		Email:  email,
		Social: social,
	}
}

func (Customer) TableName() string {
	return "film_custs"
}
