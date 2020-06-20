package models

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type Will struct {
	ClientID string          `gorm:"primary_key"`
	QoS      uint8           `gorm:"not null`
	Message  *postgres.Jsonb `gorm:"not null"`
	Topic    string          `gorm:"not null"`
}

type willGorm struct {
	db *gorm.DB
}

func NewWillService(db *gorm.DB) WillService {
	return &willGorm{
		db,
	}
}

type WillService interface {
	Create(will *Will) error
}

func (wg *willGorm) Create(will *Will) error {
	return wg.db.Create(will).Error
}
