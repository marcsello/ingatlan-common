package model

import (
	"time"
)

type SearchQuery struct {
	ID           uint      `json:"id" gorm:"primarykey,unique"`
	Source       string    `json:"src" gorm:"type:varchar(3)"`
	CreatedAt    time.Time `json:"created_at" gorm:"default:now()"`
	LastExecuted time.Time `json:"last_executed"`

	// Freetext for simpler identifying of the query
	Note string `json:"note" gorm:"type:varchar(2048)"`

	// Query is the site specific query string. It's usage and format depends on the site.
	Query string `json:"query" gorm:"type:varchar(2048)"`

	// Silent means that we don't want notifications for this query. Useful for using it only for data collection.
	Silent *bool `json:"silent" gorm:"default:false"`

	// Cooldown is in seconds: AT LEAST this much time have to be spent between running of this query
	Cooldown uint32 `json:"cooldown" gorm:"default:300"` // <- default: 5m

	Hits []*Albi `json:"hits,omitempty" gorm:"many2many:query_albis;"`
}

func (sq *SearchQuery) IsSilent() bool {
	return sq.Silent != nil && *sq.Silent
}

// IsReady in other words, is cooldown expired
func (sq *SearchQuery) IsReady(currentTime time.Time) bool {
	return sq.LastExecuted.Add(time.Duration(sq.Cooldown) * time.Second).Before(currentTime)
}
