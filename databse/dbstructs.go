package databse

import "gorm.io/gorm"

type MessageBox struct {
	gorm.Model
	Message string
}
