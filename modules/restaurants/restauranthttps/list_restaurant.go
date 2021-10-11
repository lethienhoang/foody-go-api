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

func ListRestaurantPath(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, common.NewFailureResponse(http.StatusBadRequest, err.Error()))
			return
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, common.NewFailureResponse(http.StatusBadRequest, err.Error()))
			return
		}

		store := restaurantrepo.NewSqlConn(db)

		service := restaurantservices.NewListRestaurantService(store)
		results, err := service.ListRestaurant(c.Request.Context(), nil, &filter, &paging)
		if err != nil {
			c.JSON(http.StatusInternalServerError, common.NewFailureResponse(http.StatusInternalServerError, err.Error()))
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponseWithPaging(results, paging, filter))
	}
}

