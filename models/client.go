package models

type Client struct{
	Id uint `json:"id" gorm:"primary_Key"`
	Name string `json:"name" gorm:"not null"`
	Username string `json:"username" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
	Email string `json:"email" gorm:"not null"`
	Manageclient []UserProducts `json:"manageclient" gorm:"foreignKey:ClientId"`
}