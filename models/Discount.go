package models

import ( "github.com/jinzhu/gorm" )

type Discount struct{
    gorm.Model
    Title        string  'gorm:"size:255"'
    Description  string
    SC_id        int    //'gorm:"foreignkey:Service_centre;association_foreignkey:Id"'
    Auto_id      int    //'gorm:"foreignkey:Auto;association_foreignkey:Id"'
}

//db.Model(&Discount).Related(&profile)