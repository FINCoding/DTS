package models

type ServiceCentres struct{
    id string
    adress string
    id_car string
}

func NewSC(id, adress, id_car string){
    return &ServiceCentres(id, adress, id_car)
}