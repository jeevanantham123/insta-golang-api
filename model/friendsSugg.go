package model

import "github.com/jinzhu/gorm"

//FriendSuggestion table
type FriendSuggestion struct {
	gorm.Model

	UserID uint

	UserName string

	Followed bool
}
