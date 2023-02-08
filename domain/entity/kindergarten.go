package entity

import (
	"gorm.io/gorm"
	"time"
)

type Kindergarten struct {
	ID               string         `gorm:"primary_key;column:id;type:VARCHAR;" json:"id"`
	Name             string         `gorm:"column:name;type:VARCHAR;" json:"name"`
	Address          string         `gorm:"column:address;type:VARCHAR;" json:"address"`
	NumberOfChildren int64          `gorm:"column:number_of_children;type:INT8;default:0;" json:"numberOfChildren"`
	AgeOfGirl        int64          `gorm:"column:age_of_girl;type:INT8;default:0;" json:"ageOfGirl"`
	AgeOfBoy         int64          `gorm:"column:age_of_boy;type:INT8;default:0;" json:"ageOfBoy"`
	CreatedAt        time.Time      `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"createdAt"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;type:TIMESTAMPTZ;" json:"deletedAt"`
}
