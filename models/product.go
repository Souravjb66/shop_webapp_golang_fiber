package models

type Products struct{
	Id uint `json:"id" gorm:"primary_Key"`
	Name string `json:"name" gorm:"not null"`
	ProductType string `json:"producttype" gorm:"not null"`
	Price uint `json:"price" gorm:"not null"`
	ManageProduct []UserProducts `json:"manageproduct" gorm:"foreignKey:ProductId"`
}