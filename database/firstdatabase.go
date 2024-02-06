package database

import (
	"log"
	// "os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"thirdweb/models"
)
type DbInstance struct{
	Db *gorm.DB
}

var MyDatabase DbInstance
func ConnectDB(){
	DbConnection:="root:sourav@90###@tcp(localhost:3306)/Shop?parseTime=true"
	db,err:=gorm.Open(mysql.Open(DbConnection),&gorm.Config{})
	if err!=nil{
		log.Println("error in database  :",err)
		// os.Exit(2)
		
	}
	MyDatabase.Db=db
	err=db.AutoMigrate(&models.Client{},&models.Products{},&models.UserProducts{})  //give table name first which didnt have foreign key
	if err!=nil{
		log.Println("error in migrate :",err)
		
	}
}
