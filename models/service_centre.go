package models

type Service_centre struct{
    Id string
    Adress string
}

func NewSC(Id, Adress string) *Service_centre {
    return &Service_centre{Id, Adress}
}
