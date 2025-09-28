package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	Username  string         `json:"username" gorm:"uniqueIndex;not null;size:100"`
	Email     string         `json:"email" gorm:"uniqueIndex;not null;size:255"`
	Password  string         `json:"-" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// 关联关系
	Posts    []Post    `json:"posts,omitempty" gorm:"foreignKey:UserID"`
	Comments []Comment `json:"comments,omitempty" gorm:"foreignKey:UserID"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) { return u.HashPassword() }
