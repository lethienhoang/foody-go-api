package restaurantrepo

import (
	"context"
	"github.com/foody-go-api/common"
	"github.com/foody-go-api/modules/restaurants/restaurantmodel"
)

func (s *sqlConn) ListByCondition(ctx context.Context,
	condition map[string]interface{},
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]restaurantmodel.Restaurant, error) {

	var results []restaurantmodel.Restaurant

	for i := range moreKeys {
		s.db.Preload(moreKeys[i])
	}

	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where(condition).Error; err != nil {
		return nil, err
	}

	if f := filter; f != nil {
		// filter here
	}

	if err := s.db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := s.db.Offset((paging.Page - 1) * paging.Limit).Limit(paging.Limit).Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}
