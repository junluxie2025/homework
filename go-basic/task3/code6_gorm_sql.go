package main

import (
	"fmt"

	"gorm.io/gorm"
)

func RunGrom2(db *gorm.DB) {
	//1、编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	var user User
	db.Preload("Post").Preload("Post.Comment").Where("name = ?", "Kelly").First(&user)
	fmt.Println(user)

	//2、编写Go代码，使用Gorm查询评论数量最多的文章信息。
	var post_id int
	var cnt int
	row := db.Raw("SELECT post_id,count(1) ct FROM comments group by post_id order by ct desc limit 1 ").Row()
	row.Scan(&post_id, &cnt)
	fmt.Println(post_id, cnt)
	var post Post
	db.Where("id = ?", post_id).Find(&post)
	fmt.Println(post)
}
