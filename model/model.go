package model

import (
	"database/sql"
	"time"
)

type OldPrice struct {
	ID uint `json:"id" gorm:"primarykey"`

	AlbiID     uint          `json:"albi_id"`
	AlbiSource string        `json:"albi_src" gorm:"type:varchar(3)"`
	Timestamp  time.Time     `json:"ts"`
	Price      sql.NullInt64 `json:"price"`

	Albi *Albi `json:"albi" gorm:"belongsTo:Albi"`
}

type Albi struct {
	ID           uint        `json:"id" gorm:"primarykey"`
	Source       string      `json:"src" gorm:"type:varchar(3);primarykey"`
	CreatedAt    time.Time   `json:"created_at" gorm:"default:now()"`
	LastSeen     time.Time   `json:"last_seen" gorm:"default:now()"`
	PriceHistory []*OldPrice `json:"price_history" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Price        sql.NullInt64 `json:"price"`
	Addr         string        `json:"addr" gorm:"type:varchar(200)"`
	Size         string        `json:"size" gorm:"type:varchar(100)"`
	Rooms        string        `json:"rooms" gorm:"type:varchar(100)"`
	URL          string        `json:"url" gorm:"type:varchar(255)"`
	ThumbnailURL string        `json:"thumbnail_url" gorm:"type:varchar(255)"`

	Originator []*SearchQuery `json:"originators" gorm:"many2many:query_albis;"`
}

type SearchQuery struct {
	ID           uint      `json:"id" gorm:"primarykey,unique"`
	Source       string    `json:"src" gorm:"type:varchar(3)"`
	CreatedAt    time.Time `json:"created_at" gorm:"default:now()"`
	LastExecuted time.Time `json:"last_executed"`

	Note  string `json:"note" gorm:"type:varchar(2048)"`
	Query string `json:"query" gorm:"type:varchar(2048)"`

	Hits []*Albi `json:"hits" gorm:"many2many:query_albis;"`
}

type User struct {
	// All these data are received from Telegram
	ID        uint64 `json:"id" gorm:"primarykey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	PhotoUrl  string `json:"photo_url"`
	// Those are needs to be boolean, otherwise goorm won't update them: https://stackoverflow.com/questions/56653423/gorm-doesnt-update-boolean-field-to-false
	Active	  *bool  `json:"active" gorm:"default:false"`
	Admin	  *bool  `json:"admin" gorm:"default:false"`
}
