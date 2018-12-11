package models

type Discounts struct{
    id string
    description string
    is_sc string
    id_car string
}

func NewDiscounts(id, description, id_sc, id_car string){
    return &Discounts(id, description, id_sc, id_car)
}