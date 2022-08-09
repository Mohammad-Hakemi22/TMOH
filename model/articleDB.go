package model

type Article struct {
	Id    int
	Title string
	Text  string
	Date  string
	Rate  float64
	Athor *Athor
	VIP   bool
}