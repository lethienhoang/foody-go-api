package restauranthttps

import (
	"github.com/foody-go-api/common"
	"github.com/foody-go-api/modules/restaurants/restaurantmodel"
	"github.com/foody-go-api/modules/restaurants/restaurantrepo"
	"github.com/foody-go-api/modules/restaurants/restaurantservices"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateRestaurantPath(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.NewFailureResponse(http.StatusBadRequest, err.Error()))
			return
		}

		store := restaurantrepo.NewSqlConn(db)

		service := restaurantservices.NewCreateRestaurantService(store)
		if err := service.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusInternalServerError, common.NewFailureResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponseNoPaging(&data))
	}
}

