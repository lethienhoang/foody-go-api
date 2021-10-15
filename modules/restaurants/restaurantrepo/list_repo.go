package restaurantrepo

import (
	"context"
	"github.com/foody-go-api/common"
	"github.com/foody-go-api/modules/restaurants/restaurantmodel"
)

func (s *SqlConn) ListByCondition(ctx context.Context,
	condition map[string]interface{},
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]restaurantmodel.Restaurant, error) {

	var results []restaurantmodel.Restaurant

	for i := range moreKeys {
		s.db.Preload(moreKeys[i])
	}

	s.db = s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where(condition).Where("status in (1)")

	if f := filter; f != nil {
		// filter here
	}

	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&results).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return results, nil
}
