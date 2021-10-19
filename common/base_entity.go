package common

import "time"

type BaseEntity struct {
	Id        int        `json:"id" gorm:"column:id;unique"`
	Status    int        `json:"status" gorm:"status"`
	CreatedAt *time.Time `json:"created_at" gorm:"created_at"`
	UpdateAt  *time.Time `json:"update_at" gorm:"update_at"`
}
