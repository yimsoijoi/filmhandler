package main

import (
	"encoding/json"
	"fmt"

	"github.com/yimsoijoi/filmhandler/datamodel"
	"github.com/yimsoijoi/filmhandler/lib/postgres"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB, conf postgres.Config) {
	if conf.Migrate {
		c := datamodel.Customer{}
		o := datamodel.Order{}
		db.AutoMigrate(&c, &o)
	}
}

func check(db *gorm.DB) {
	c1 := datamodel.NewCustomer("pak", "0614397379", "p.winaitammakul@gmail.com", "ppakzv")
	o1 := datamodel.NewOrder(c1.Tel, 69, "kodakColor200", "colorNeg", "takeOrders", true)
	_c1, _ := json.MarshalIndent(c1, "  ", "  ")
	_o1, _ := json.MarshalIndent(o1, "  ", "  ")
	fmt.Printf("{\n  \"c1\": %s,\n  \"01\": %s,\n}", _c1, _o1)

	db.Create(c1)
	db.Create(o1)

	var dbOrder datamodel.Order

	db.Preload("Customer").First(&dbOrder)

	_dbOrder, _ := json.MarshalIndent(dbOrder, "  ", "  ")
	_dbCust, _ := json.MarshalIndent(dbOrder.Customer, " ", " ")
	fmt.Printf("%s\n%s\n", _dbOrder, _dbCust)

}
