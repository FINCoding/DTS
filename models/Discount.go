package models

import ( "github.com/jinzhu/gorm" )

type Discount struct{
    gorm.Model
    Title        string  'gorm:"size:255"'
    Description  string
    SC_id        int
    Auto_id      int
}

func NewDiscount(id, description, SC_id, Auto_id string) *Discount {
    return &Discount{id, description, SC_id, Auto_id}

db.Model(&Discount).Related(&profile)
