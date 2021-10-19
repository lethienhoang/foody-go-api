package uploadfilerepo

import (
	"context"
	"github.com/foody-go-api/common"
)

func (store *SqlConn) DeleteImages(ctx context.Context, ids []int) error {
	db := store.db

	if err := db.Table(common.Image{}.TableName()).
		Where("id in (?)", ids).
		Delete(nil).Error; err != nil {
		return err
	}

	return nil
}
