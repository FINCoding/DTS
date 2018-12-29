package models

/*
type Service_centres struct{
    gorm.Model
    Id int          `gorm:"primary_key"`
    Adress Adress
    IdAuto string
}

type Address struct {
    gorm.Model
    ID       int            `gorm:"primary_key"`
    Country  string         `gorm:"not null"` // Set field as not nullable and unique
    State    string         `gorm:"type:varchar(100)"`
    City     string         `gorm:"not null"`
    Street   string
    House    string
    Room     string
    Other    string
}

db.Model(&service_centres).Related(&address)

type auto struct{
    Id int
    Brand string
    Model string
    IdSc string
}

type Discount struct{
    gorm.Model

    Id int             `gorm:"primary_key"`
    Description string `gorm:"size:255"`
    IdSC int
    IdAuto int
}
*/