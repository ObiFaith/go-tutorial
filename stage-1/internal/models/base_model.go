package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now().UTC()

	b.ID = uuid.Must(uuid.NewV7())
	b.CreatedAt = now
	b.UpdatedAt = now

	return
}

func (b *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	b.UpdatedAt = time.Now().UTC()
	return
}
