package restauranthttps

import (
	"github.com/foody-go-api/common"
	"github.com/foody-go-api/modules/restaurants/restaurantmodel"
	"github.com/foody-go-api/modules/restaurants/restaurantrepo"
	"github.com/foody-go-api/modules/restaurants/restaurantservices"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func UpdateRestaurantPath(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.NewFailureResponse(http.StatusBadRequest, err.Error()))
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.NewFailureResponse(http.StatusBadRequest, err.Error()))
		}

		store := restaurantrepo.NewSqlConn(db)

		service := restaurantservices.NewUpdateRestaurantService(store)
		if err := service.Update(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusInternalServerError, common.NewFailureResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponseNoPaging(&data))
	}
}

