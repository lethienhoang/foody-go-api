package restaurantrepo

import (
	"context"
	"github.com/foody-go-api/common"
	"github.com/foody-go-api/modules/restaurants/restaurantmodel"
	"gorm.io/gorm"
)

func (s *SqlConn) FindByCondition(ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string) (*restaurantmodel.Restaurant, error) {

	var result restaurantmodel.Restaurant

	for i := range moreKeys {
		s.db.Preload(moreKeys[i])
	}

	s.db = s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where(condition)

	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrEntityNotFound(restaurantmodel.Restaurant{}.TableName(), common.RecordNotFound)
		}

		return nil, common.ErrDB(err)
	}

	return &result, nil
}
