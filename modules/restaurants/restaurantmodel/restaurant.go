package restaurantmodel

import (
	"errors"
	"github.com/foody-go-api/common"
	"strings"
)

type Restaurant struct {
	common.BaseEntity `json:",inline"`
	//Id   int    `json:"id" gorm:"column:id;unique"`
	Name   string `json:"name" gorm:"column:name;"`
	Addr   string `json:"address" gorm:"column:addr;"`
	Status int    `json:"status" gorm:"column:status;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name   *string `json:"name" gorm:"column:name;"`
	Addr   *string `json:"address" gorm:"column:addr;"`
	Status int     `json:"status" gorm:"column:status;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	common.BaseEntity `json:",inline"`
	Name   string `json:"name" gorm:"column:name;"`
	Addr   string `json:"address" gorm:"column:addr;"`
	Status int    `json:"status" gorm:"column:status;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantCreate) validate() error {
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return errors.New("restaurant name can't be blank")
	}

	return nil
}
