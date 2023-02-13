package entity

import (
	"gorm.io/gorm"
	"time"
)

type Kindergarten struct {
	ID               string         `gorm:"primary_key;column:kindergarten_id;type:VARCHAR;" json:"kindergartenId"`
	Name             string         `gorm:"column:name;type:VARCHAR;" json:"name"`
	NumberOfTeacher  string         `gorm:"column:number_of_teacher;type:VARCHAR;" json:"number_of_teacher"`
	Area             string         `gorm:"column:area;type:VARCHAR;" json:"area"`
	Address          string         `gorm:"column:address;type:VARCHAR;" json:"address"`
	Personnel        int64          `gorm:"column:personnel;type:INT8;default:0;" json:"personnel"`
	NumberOfChildren int64          `gorm:"column:number_of_children;type:INT8;default:0;" json:"numberOfChildren"`
	UrlOfHomepage    string         `gorm:"column:url_of_homepage;type:VARCHAR;" json:"urlOfHomepage"`
	Introduction     string         `gorm:"column:introduction;type:VARCHAR;" json:"introduction"`
	PhoneNumber      string         `gorm:"column:phone_number;type:VARCHAR;" json:"phoneNumber"`
	CreatedAt        time.Time      `gorm:"column:created_at;type:TIMESTAMPTZ;default:CURRENT_TIMESTAMP;" json:"createdAt"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;type:TIMESTAMPTZ;" json:"deletedAt"`
}

// 외부에서 사용하는 구조체는 대문자로 시작
