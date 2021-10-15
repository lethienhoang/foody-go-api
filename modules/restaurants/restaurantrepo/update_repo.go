package restaurantrepo

import (
	"context"
	"github.com/foody-go-api/common"
	"github.com/foody-go-api/modules/restaurants/restaurantmodel"
)

func (s *SqlConn) Update(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func (s *SqlConn) Delete(ctx context.Context, id int) error {
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
		"status": 0,
	}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
