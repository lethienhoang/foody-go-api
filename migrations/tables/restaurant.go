package tables

import (
	"fmt"
	"github.com/foody-go-api/modules/restaurants/restaurantmodel"
	"gorm.io/gorm"
	"log"
)

func MigrationRestaurantTable(db *gorm.DB)  {
	hasTable := db.Migrator().HasTable(&restaurantmodel.Restaurant{})
	if !hasTable {
		err := db.Migrator().CreateTable(&restaurantmodel.Restaurant{})
		if err != nil {
			fmt.Printf("migration is not working - restaurant table - error: %s", err.Error())
			log.Fatalf("migration is not working - restaurant table - error: %s", err.Error())
		}

		//err = db.Migrator().CreateIndex(&restaurantmodel.Restaurant{}, "Id")
		//if err != nil {
		//	fmt.Printf("migration can not create index - error: %s", err.Error())
		//	log.Panicf("migration can not create index - error: %s", err.Error())
		//}
		//
		//err = db.Migrator().CreateIndex(&restaurantmodel.Restaurant{}, "idx_id")
		//if err != nil {
		//	fmt.Printf("migration can not create index - error: %s", err.Error())
		//	log.Panicf("migration can not create index - error: %s", err.Error())
		//}
	}
}
