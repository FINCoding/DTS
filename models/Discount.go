package models

import ( "github.com/jinzhu/gorm" )

type Discount struct{
    gorm.Model
    Title        string  'gorm:"size:255"'
    Description  string
    SC_id        int
    Auto_id      int
}

db.Model(&Discount).Related(&profile)