package model


type Article struct {
	Id       int     `gorm:"primaryKey"`
	Title    string  `json:"title"`
	Text     string  `json:"text"`
	Date     string  `json:"date"`
	Rate     float64 `json:"rate"`
	Vip      bool    `json:"vip"`
	AuthorID int     `gorm:"foreignkey"`
}
