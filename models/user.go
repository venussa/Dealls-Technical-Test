package models

import (
  gorm "gorm.io/gorm"
)

type User struct {
  gorm.Model
  Email string `gorm:"type:varchar(255);uniqueIndex;not null"`
  FullName string `gorm:"type:varchar(255);not null"`
  ProfilePicture string `gorm:"type:text;not null"`
  IsVerified int `gorm:"type:int;default 0"`
  Password string `gorm:"type:varchar(255);not null"`
}