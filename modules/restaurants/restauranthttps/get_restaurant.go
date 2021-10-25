package restauranthttps

import (
	"github.com/foody-go-api/common"
	"github.com/foody-go-api/modules/restaurants/restaurantrepo"
	"github.com/foody-go-api/modules/restaurants/restaurantservices"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetByIdRestaurantPath(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		crypto:= common.NewCrypto()
		crypto.Decrypt(c.Param("id"), common.DbTypeRestaurant)

		id, err := strconv.Atoi(crypto.ID)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		conn := restaurantrepo.NewSqlConn(db)

		service := restaurantservices.NewFindRestaurantService(conn)
		data, err := service.FindByCondition(c.Request.Context(), map[string]interface{}{"id": id})
		if err != nil {
			if err == common.RecordNotFound {
				panic(common.ErrNotFound(err))
			}

			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponseNoPaging(&data))
	}
}
