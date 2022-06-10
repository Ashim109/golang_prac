package main

import (
	"fmt"
	"net/http"



	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	CustomerName  string `json:"customername"`
	CustomerEmail string `json:"customeremail"`
}
type Payment struct {
	gorm.Model
	Payments      int64  `json:"payments"`
	CustomerEmail string `json:"customeremail"`
}

//   *gorm.DB
func Init() {
	// dbURL := "postgres://postgres:aditya1@localhost:5432/customer"
	dbURL := "host=localhost user=postgres password=Simple@1 dbname=customer port=5432 sslmode=disable"
	Database, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	Database.AutoMigrate(&User{})
	Database.AutoMigrate(&Payment{})
	DB = Database
	// return DB
}
func GetUsers(c *gin.Context) {
	var users []User

	DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func Person(c *gin.Context) {
	var users User

	if err := DB.Where("customer_name=?", c.Param("customer_name")).First(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
	}
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func main() {

	Init()

	r := gin.Default()

	r.GET("/users", GetUsers)
	r.GET("/users/:customer_name", Person)
	r.Run(":9000")

}
