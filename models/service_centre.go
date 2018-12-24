package models

type Service_centre struct{
    Id int
    Adress Adress
}

type Adress struct {
    Id       int
    Country  string
    State    string
    City     string
    Street   string
    House    string
    Room     string
    Other    string
}
