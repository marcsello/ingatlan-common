package model

import (
	"database/sql"
	"time"
)

type OldPrice struct {
	ID uint `gorm:"primarykey"`

	AlbiID     uint
	AlbiSource string `gorm:"type:varchar(3)"`
	Timestamp  time.Time
	Price      sql.NullInt64

	Albi *Albi `gorm:"belongsTo:Albi"`
}

type Albi struct {
	ID           uint        `gorm:"primarykey"`
	Source       string      `gorm:"type:varchar(3);primarykey"`
	CreatedAt    time.Time   `gorm:"default:now()"`
	LastSeen     time.Time   `gorm:"default:now()"`
	PriceHistory []*OldPrice `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Price        sql.NullInt64
	Addr         string `gorm:"type:varchar(200)"`
	Size         string `gorm:"type:varchar(100)"`
	Rooms        string `gorm:"type:varchar(100)"`
	URL          string `gorm:"type:varchar(255)"`
	ThumbnailURL string `gorm:"type:varchar(255)"`

	Originator []*SearchQuery `gorm:"many2many:query_albis;"`
}

type SearchQuery struct {
	ID           uint      `gorm:"primarykey,unique"`
	Source       string    `gorm:"type:varchar(3)"`
	CreatedAt    time.Time `gorm:"default:now()"`
	LastExecuted time.Time

	Note  string `gorm:"type:varchar(2048)"`
	Query string `gorm:"type:varchar(2048)"`

	Albis []*Albi `gorm:"many2many:query_albis;"`
}
