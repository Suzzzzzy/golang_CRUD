package entity

import (
	"gorm.io/gorm"
	"time"
)

type TeacherInfo struct {
	ID           string         `gorm:"primary_key;column:id;type:VARCHAR;" json:"id"`
	Name         string         `gorm:"column:name;type:VARCHAR;" json:"name"`
	ProfileUrl   string         `gorm:"column:profile_url;type:VARCHAR;" json:"profileUrl"`
	Age          int64          `gorm:"column:age;type:INT8;" json:"age"`
	Introduction string         `gorm:"column:introduction;type:VARCHAR;" json:"introduction"`
	Major        string         `gorm:"column:major;type:VARCHAR;" json:"major"`
	Certificate  string         `gorm:"column:certificate;type:VARCHAR;" json:"certificate"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"createdAt"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;type:TIMESTAMPTZ;" json:"deletedAt"`
}
