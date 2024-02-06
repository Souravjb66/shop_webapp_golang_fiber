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
		Clientname string `json:"clientname" gorm:"not null"`
		Username string `json:"username" gorm:"not null"`
		Password string `json:"password" gorm:"not null"`
		Email string `json:"email" gorm:"not null"`
		Productname string `json:"productname" gorm:"not null"`
		ProductType string `json:"producttype" gorm:"not null"`
		Price uint `json:"price" gorm:"not null"`

	}
	var myuser User
	err:=c.BodyParser(&myuser)
	if err!=nil{
		log.Println("error in create client",err)
	}
	myclient:=models.Client{
		Name:myuser.Clientname,
		Username: myuser.Username,
		Password: myuser.Password,
		Email: myuser.Email,
	}
	myproduct:=models.Products{
		Name:myuser.Productname,
		ProductType: myuser.ProductType,
		Price:myuser.Price,

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
	var client []models.Client
	var product []models.Products

	one:=database.MyDatabase.Db.Find(&client)
	if one!=nil{
		log.Println(one)
	}
	two:=database.MyDatabase.Db.Find(&product)
	if two!=nil{
		log.Println(two)
	}
	values:=[]AllValues{}
	for _,user:=range client{
		values=append(values,AllValues{
			ClientId:user.Id,
			ClientName:user.Name,
			ClientUsername:user.Username,
			ClientPassword:user.Password,
			ClientEmail:user.Email,
		
		} )


	}
	for _,item:=range product{
		values=append(values, AllValues{
			ProductId:item.Id,
				ProductName:item.Name,
				ProductType:item.ProductType,
				ProductPrice: item.Price,

		})
	}

	
	return c.Status(200).JSON(values)

}
