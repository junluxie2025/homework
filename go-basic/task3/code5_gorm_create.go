package main

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
	Post []Post
}

type Post struct {
	gorm.Model
	UserID  int
	Title   string
	Comment []Comment
}

type Comment struct {
	gorm.Model
	PostID  int
	Content string
}

func RunGrom1(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
	insert(db)
}

func insert(db *gorm.DB) {
	users := []User{
		{Name: "Alice",
			Age: 20,
			Post: []Post{
				{
					UserID: 1,
					Title:  "第一篇文章",
					Comment: []Comment{
						{PostID: 1, Content: "第一条评论"},
						{PostID: 1, Content: "第二条评论"},
					},
				},
				{
					UserID: 1,
					Title:  "第二篇文章",
					Comment: []Comment{
						{PostID: 1, Content: "第一条评论"},
						{PostID: 1, Content: "第二条评论"},
						{PostID: 1, Content: "第三条评论"},
					},
				},
			},
		},
		{
			Name: "Kelly",
			Age:  20,
			Post: []Post{
				{

					Title: "第一篇文章",
					Comment: []Comment{
						{Content: "第一条评论"},
					},
				},
				{

					Title: "第二篇文章",
					Comment: []Comment{
						{Content: "第一条评论"},
					},
				},
			},
		},
	}
	db.Create(&users)
}
