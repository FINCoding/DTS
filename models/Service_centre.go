package models

type Service_centre struct{
    Id string
    Adress string
    IdAuto string
}

func NewSC(Id, Adress, IdAuto string) *Service_centre {
    return &Service_centre{Id, Adress, IdAuto}
}
