package service

import (
	"log"
	"thirdweb/database"
	"thirdweb/models"
	"github.com/gofiber/fiber/v2"
)

func GetIndex(c *fiber.Ctx)error{
	return c.Render("frontend/index.html",fiber.Map{})

}
func CreateClient(c *fiber.Ctx)error{
	type User struct{
		Name string `json:"name" gorm:"not null"`
		Username string `json:"username" gorm:"not null"`
		Password string `json:"password" gorm:"not null"`
		Email string `json:"email" gorm:"not null"`
		Myproducts models.Products `json:"myproducts" gorm:"no null"`
	}
	var myuser User
	err:=c.BodyParser(&myuser)
	if err!=nil{
		log.Println("error in create client",err)
	}
	myclient:=models.Client{
		Name:myuser.Name,
		Username: myuser.Username,
		Password: myuser.Password,
		Email: myuser.Email,
	}
	myproduct:=models.Products{
		Name:myuser.Myproducts.Name,
		ProductType: myuser.Myproducts.ProductType,
		Price:myuser.Myproducts.Price,

	}
	
	output1:=database.MyDatabase.Db.Create(&myclient)
	if output1!=nil{
		log.Println(output1)
	}else{
		log.Println("error in create client product")
	}
	output2:=database.MyDatabase.Db.Create(&myproduct)
	if output2!=nil{
		log.Println(output2)
	}else{
		log.Println("erorr in creating product")
	}
	bothUserProduct:=models.UserProducts{ //we assign it after main 2 tables values are inserted after that in that name have current value id  so we can assign it in combine table
		ClientId:myclient.Id,
		ProductId:myproduct.Id,
	}
	output3:=database.MyDatabase.Db.Create(&bothUserProduct)
	if err!=nil{
		log.Println(output3)
	}
	return c.Status(200).JSON(myuser)

}
func GetAllClientsWithProduct(c *fiber.Ctx)error{
	type AllValues struct{
		ClientId uint `json:"id"`
		ClientName string `json:"clientname"`
		ClientUsername string `json:"clientusername"`
		ClientPassword string `json:"clientpassword"`
		ClientEmail string `json:"clientemail"`
		ProductId uint `json:"productid"`
		ProductName string  `json:"productname"`
		ProductType string `json:"producttype"`
		ProductPrice uint `json:"productprice"`
	}
	var client models.Client
	var product models.Products

	one:=database.MyDatabase.Db.Find(&client)
	if one!=nil{
		log.Println(one)
	}
	two:=database.MyDatabase.Db.Find(&product)
	if two!=nil{
		log.Println(two)
	}

	values:=AllValues{
		ClientId: client.Id,
		ClientName: client.Name,
		ClientUsername: client.Username,
		ClientPassword: client.Password,
		ClientEmail:client.Email,
		ProductId: product.Id,
		ProductName: product.Name,
		ProductType: product.ProductType,
		ProductPrice: product.Price,
	}
	return c.Status(200).JSON(values)

}
