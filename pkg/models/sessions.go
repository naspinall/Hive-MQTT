package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Session struct {
	ClientID    string `gorm:"primary_key"`
	Username    string
	LastConnect time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}

type sessionGorm struct {
	db *gorm.DB
}

func NewSessionService(db *gorm.DB) SessionService {
	return &sessionGorm{
		db,
	}
}

type SessionService interface {
	Create(session *Session) error
}

func (sg *sessionGorm) Create(session *Session) error {
	return sg.db.Create(session).Error
}
