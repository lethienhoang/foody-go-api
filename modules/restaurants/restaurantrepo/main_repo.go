package restaurantrepo

import "gorm.io/gorm"

type sqlConn struct {
	db *gorm.DB
}

func NewSqlConn(db *gorm.DB) * sqlConn  {
	return &sqlConn{db: db}
}
