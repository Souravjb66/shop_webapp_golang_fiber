package models

type UserProducts struct{
	Id uint `json:"id" gorm:"primary_Key"`
	ClientId uint `json:"clientid"`
	ProductId uint `json:"productid"`
}