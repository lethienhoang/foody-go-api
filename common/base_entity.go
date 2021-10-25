package common

import (
	"strconv"
	"time"
)

type BaseEntity struct {
	Id        int        `json:"-" gorm:"column:id;unique"`
	FakeId    string     `json:"id" gorm:"-"`
	Status    int        `json:"status" gorm:"status"`
	CreatedAt *time.Time `json:"created_at" gorm:"created_at"`
	UpdateAt  *time.Time `json:"update_at" gorm:"update_at"`
}

func (b *BaseEntity) EncryptID(dbType int)  {
	crypto:= NewCrypto()
	crypto.Encrypt(strconv.Itoa(b.Id), dbType)
	b.FakeId = crypto.FakeID
}
