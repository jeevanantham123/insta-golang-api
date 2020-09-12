package model

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

//UserFriend  table
type UserFriend struct {
	gorm.Model
	UserName  string
	Requested pq.StringArray `gorm:"type:varchar(64)[]"`
	Friends   pq.StringArray `gorm:"type:varchar(64)[]"`
}

//SuggestionTab structure
type SuggestionTab struct {
	ID         int    `json:"id"`
	UserName   string `json:"username"`
	Followed   bool   `json:"followed"`
	ProfileURL string `json:"profileurl"`
}
