package restaurantrepo

import (
	"context"
	"github.com/foody-go-api/common"
	"github.com/foody-go-api/modules/restaurants/restaurantmodel"
)

func (s *SqlConn) Create (ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	data.Status = 1
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
