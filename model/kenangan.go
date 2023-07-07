package model

import "time"

type Kenangan struct {
	ID          int       `json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Image       string    `gorm:"not null" json:"image"`
	UploadDate  time.Time `gorm:"not null" json:"upload_date"`
	Description string    `gorm:"not null" json:"description"`
	UserID      int       `json:"user_id"`
}

type KenanganUser struct {
	ID          int       `json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Image       string    `gorm:"not null" json:"image"`
	UploadDate  time.Time `gorm:"not null" json:"upload_date"`
	Description string    `gorm:"not null" json:"description"`
}

func (KenanganUser) TableName() string {
	return "kenangans"
}
