package datamodel

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrdUuid   string    `json:"order_uuid" gorm:"primaryKey;column:uuid;not null"`
	CustTel   string    `json:"cust_tel" gorm:"column:cust_tel;not null"`
	Customer  Customer  `json:"customer" gorm:"foreignKey:CustTel;references:Tel;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;not null"`
	FilmNo    string    `json:"film_number" gorm:"foreignKey;column:filmnumber;not null"`
	FilmName  string    `json:"film_name" gorm:"column:filmname;not null"`
	FilmType  string    `json:"film_type" gorm:"column:filmtype;not null"`
	Status    string    `json:"status" gorm:"status;not null"`
	Keep      bool      `json:"keep" gorm:"keep;not null"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`
	//ลิ้งค์ primary key มาตัวเดียวพอ
	//มี json เพื่อสำหรับให้fiber marshal ก่อนส่งกลับไป
}

func NewOrder(tel string, fnum int, fname string, ftype string, status string, keep bool) *Order {
	_fnum := fmt.Sprintf("%04d", fnum) //pad zeroes of length
	return &Order{
		OrdUuid:  uuid.NewString(),
		CustTel:  tel,
		FilmNo:   _fnum,
		FilmName: fname,
		FilmType: ftype,
		Status:   status,
		Keep:     keep,
	}
}

func (Order) TableName() string {
	return "film_orders"
}
