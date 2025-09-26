package main

import "gorm.io/gorm"

type student struct {
	gorm.Model
	Name  string
	Age   int
	Grade string
}

func Run(db *gorm.DB) {
	//db.AutoMigrate(&student{})
	student := &student{
		Name:  "张三",
		Age:   18,
		Grade: "三年级",
	}
	//编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	db.Create(student)

	//编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。

	//编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	//编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
}
