package context

import (
	"fmt"
	"github.com/foody-go-api/migrations/tables"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type DbCtx struct {
	DB *gorm.DB
}

func NewDbContext() *DbCtx {
	//dsn := os.Getenv("DBConnectionString")
	dsn := "root:Localhost123@tcp(127.0.0.1:3306)/foody_db?charset-utf8&parseTime=True&loc=Local"
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln(err)
	}

	// migration tables
	tables.MigrationRestaurantTable(db)

	return &DbCtx{DB: db}
}
