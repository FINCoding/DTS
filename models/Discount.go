package models

import ( "github.com/jinzhu/gorm" )

type Discount struct{
    gorm.Model
    Title        string  //'gorm:"string"'
    Description  string
    SC_id        int    //'gorm:"foreignkey:Service_centre;association_foreignkey:Id"'
    Auto_id      int    //'gorm:"foreignkey:Auto;association_foreignkey:Id"'
}


//func NewDiscount(id, description, SC_id, Auto_id string) *Discount {
//    return &Discount{id, description, SC_id, Auto_id}

//db.Model(&Discount).Related(&profile)

