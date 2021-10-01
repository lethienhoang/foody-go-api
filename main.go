package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	db *gorm.DB
)
func main() {
	runDB()

	if err := runService(); err != nil {
		log.Panicln(err)
	}
}

func runDB() *gorm.DB {
	dsn := os.Getenv("DBConnectionString")
	var err error
	//dsn := "root:Localhost123@tcp(127.0.0.1:3306)/foody_db?charset-utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln(err)
	}

	return db
}

func runService() error {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})

	return r.Run(":8081")
}
