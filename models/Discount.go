package models

type Discount struct{
    Id          string
    Description string
    IdSC        string
    IdAuto      string
}

func NewDiscount(id, description, idSC, idAuto string) *Discount {
    return &Discount{id, description, idSC, idAuto}

}