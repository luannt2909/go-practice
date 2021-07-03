package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name   string `json:"name"`
	Age    int32  `json:"age"`
	Gender Gender `json:"gender"`
}
