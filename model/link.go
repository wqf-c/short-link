package model

type Link struct {
	ID        uint   `gorm:"column:id"`
	ShortLink string `gorm:"column:shortLink"`
	LongLink  string `gorm:"column:longLink"`
}
