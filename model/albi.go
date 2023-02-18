package model

import (
	"github.com/marcsello/ingatlan-common/tools"
	"gopkg.in/guregu/null.v4"
	"time"
)

type OldPrice struct {
	ID uint `json:"id" gorm:"primarykey"`

	AlbiID     uint      `json:"albi_id"`
	AlbiSource string    `json:"albi_src" gorm:"type:varchar(3)"`
	Timestamp  time.Time `json:"ts"`
	Price      null.Int  `json:"price"`

	Albi *Albi `json:"albi,omitempty" gorm:"belongsTo:Albi"`
}

type Albi struct {
	ID           uint        `json:"id" gorm:"primarykey"`
	Source       string      `json:"src" gorm:"type:varchar(3);primarykey"`
	CreatedAt    time.Time   `json:"created_at" gorm:"default:now()"`
	LastSeen     time.Time   `json:"last_seen" gorm:"default:now()"`
	PriceHistory []*OldPrice `json:"price_history" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Price        null.Int    `json:"price"`
	Addr         string      `json:"addr" gorm:"type:varchar(200)"`
	Size         string      `json:"size" gorm:"type:varchar(100)"`
	Rooms        string      `json:"rooms" gorm:"type:varchar(100)"`
	URL          string      `json:"url" gorm:"type:varchar(255)"`
	ThumbnailURL null.String `json:"thumbnail_url" gorm:"type:varchar(255)"`

	Ignored *bool `json:"ignored" gorm:"default:false"`

	Originators []*SearchQuery `json:"originators,omitempty" gorm:"many2many:query_albis;"`
}

func (a *Albi) IsIgnored() bool {
	return a.Ignored != nil && *a.Ignored
}

// HaveValidThumbnail uses tools.IsValidImageUrl to determine if the Albi have a valid thumbnail url. (It will make a HTTP request!)
func (a *Albi) HaveValidThumbnail() bool {

	if !a.ThumbnailURL.Valid {
		return false
	}

	return tools.IsValidImageUrl(a.ThumbnailURL.String)
}
