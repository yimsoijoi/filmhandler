package datamodel

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrdUuid string `json:"order_uuid" gorm:"primaryKey;column:uuid;not null"`
	// CustUuidInfo Info      `json:"uuidinfo" gorm:"column:customer_uuid;not null"`
	// CustUuid     Customer  `json:"customer_uuid" gorm:"foreignKey:CustUuidInfo;references:Address;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;not null"`
	NameInfo  Info      `json:"name" gorm:"column:name;not null"` //why use json?
	Name      Customer  `json:"-" gorm:"foreignKey:NameInfo;references:Address;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;not null"`
	TelInfo   Info      `json:"tel" gorm:"column:tel;not null"`
	Tel       Customer  `json:"-" gorm:"foreignKey:TelInfo;references:Address;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;not null"`
	EmailInfo Info      `json:"e-mail" gorm:"column:e-mail"`
	Email     Customer  `json:"-" gorm:"foreignKey:EmailInfo;references:Address;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;not null"`
	FilmNo    int       `json:"filmnumber" gorm:"foreignKey;column:filmnumber;not null"`
	FilmName  string    `json:"filmname" gorm:"column:filmname;not null"`
	FilmType  string    `json:"filmtype" gorm:"column:filmtype;not null"`
	Status    string    `json:"status" gorm:"status;not null"`
	Keep      bool      `json:"keep" gorm:"keep;not null"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`
}

func NewOrder(name Customer, tel Customer, email Customer, fnum int, fname string, ftype string, status string, keep bool) *Order {
	return &Order{
		OrdUuid:  uuid.NewString(),
		Name:     name,
		Tel:      tel,
		Email:    email,
		FilmNo:   fnum,
		FilmName: fname,
		FilmType: ftype,
		Status:   status,
		Keep:     keep,
	}
}
