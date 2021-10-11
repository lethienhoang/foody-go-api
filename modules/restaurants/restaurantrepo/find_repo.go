package restaurantrepo

import (
	"context"
	"github.com/foody-go-api/modules/restaurants/restaurantmodel"
)

func (s *sqlConn) FindByCondition(ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string) (*restaurantmodel.Restaurant, error) {

	var result restaurantmodel.Restaurant

	for i := range moreKeys {
		s.db.Preload(moreKeys[i])
	}

	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where(condition).Error; err != nil {
		return nil, err
	}

	if err := s.db.First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}
