package db

import (
	"time"
)

type Article struct {
	Id    int
	Title string
	Text  string
	Date  time.Time
	Rate  float64
	Athor *Athor
}

type Athor struct {
	Name string
	Bio  string
	Age  int
}

