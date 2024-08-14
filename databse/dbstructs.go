package databse

import "gorm.io/gorm"

type MessageBox struct {
	gorm.Model
	Message string
	Box     int64
}


type SearchResult[T any] struct {
	Response *T
	Found bool
}