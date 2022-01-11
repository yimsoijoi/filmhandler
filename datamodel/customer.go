package datamodel

import (
	"time"

	"github.com/google/uuid"
)

type Info string
type Customer struct {
	Uuid      string    `json:"uuid" gorm:"primaryKey;column:uuid;not null"`
	Name      Info      `json:"name" gorm:"column:name;not null"`
	Tel       Info      `json:"tel" gorm:"secondaryKey;column:tel;not null"`
	Email     Info      `json:"email" gorm:"column:email;not null"`
	LineId    Info      `json:"lineId" gorm:"column:lineId"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`
}

func NewCustomer(name Info, tel Info, email Info, lineId Info) *Customer {
	return &Customer{
		Uuid:      uuid.NewString(),
		Name:      name,
		Tel:       tel,
		Email:     email,
		LineId:    lineId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
