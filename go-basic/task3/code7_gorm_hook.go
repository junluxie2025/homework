package main

import (
	"fmt"

	"gorm.io/gorm"
)

func RunGrom3(db *gorm.DB) {

}

func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Parent BeforeCreate")
	return nil
}

func (c *Comment) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Println("Child BeforeCreate")
	return nil
}
