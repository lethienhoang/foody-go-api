package routes

import (
	"github.com/foody-go-api/modules/restaurants/restauranthttps"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RestaurantRoute(r *gin.Engine, db *gorm.DB)  {
	restaurant := r.Group("/restaurants")
	{
		restaurant.POST("", restauranthttps.CreateRestaurantPath(db))
		restaurant.GET("", restauranthttps.ListRestaurantPath(db))
		restaurant.PUT("", restauranthttps.UpdateRestaurantPath(db))
	}
}
