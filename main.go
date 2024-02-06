package main

import (
	// "fmt"
	"log"
	"thirdweb/database"
	"thirdweb/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	
)

func main(){
	database.ConnectDB()
	app:=fiber.New()
	app.Use(cors.New())
	routes(app)

	defer func(){
		mydb,_:=database.MyDatabase.Db.DB()
		err:=mydb.Close()
		if err!=nil{
			log.Println("main f",err)
		}
	}()
	log.Fatal(app.Listen(":8080"))

}
func routes(app *fiber.App){
	app.Get("/",service.GetIndex)
}