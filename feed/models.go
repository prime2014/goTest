package feed

import (
	"accounts"
	"time"
)

type Feed struct {
	ID       uint            `json:"id"`
	Author   *accounts.Users `json:"author"`
	AuthorID uint            `json:"author_id"`
	Post     string          `json:"post"`
	Medias   []Media         `json:"medias"`
	Likes    uint            `json:"likes" gorm:"default:0"`
	Views    uint            `json:"views" gorm:"default:0"`
	PubDate  time.Time       `json:"pubdate" gorm:"default:current_timestamp"`
}

type Media struct {
	ID       uint      `json:"id"`
	FeedID   uint      `json:"feed_id"`
	FilePath string    `json:"filepath"`
	Width    float64   `json:"width"`
	Height   float64   `json:"height"`
	PubDate  time.Time `json:"pubdate" gorm:"default:current_timestamp"`
}
