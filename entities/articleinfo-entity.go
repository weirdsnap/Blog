package entities

import (
//"time"
)

// ArticleInfo .
type ArticleInfo struct {
	AID     int `orm:"id,auto-inc"` //语义标签
	Title   string
	Class   string
	Content string
}

// NewArticleInfo .
func NewArticleInfo(a ArticleInfo) *ArticleInfo {
	if len(a.Title) == 0 {
		panic("Title shold not null!")
	}
	return &a
}
