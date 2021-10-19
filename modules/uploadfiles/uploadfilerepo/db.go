package uploadfilerepo

import "gorm.io/gorm"

type SqlConn struct {
	db *gorm.DB
}

func NewSqlConn(db *gorm.DB) *SqlConn {
	return &SqlConn{db: db}
}
