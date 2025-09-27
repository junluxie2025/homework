package main

import (
	"gorm.io/gorm"
)

type student struct {
	gorm.Model
	Name  string
	Age   int
	Grade string
}

func RunCRUD(db *gorm.DB) {
	//db.AutoMigrate(&student{})

	//students := []*student{
	//	{Name: "张三", Age: 18, Grade: "三年级"},
	//	{Name: "李四", Age: 8, Grade: "四年级"},
	//	{Name: "王五", Age: 20, Grade: "五年级"},
	//}
	//编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	//db.Debug().Create(students)

	//编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	//var student student
	//db.Where("age > ?", "18").First(&student)
	//fmt.Println(student)

	//编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	//var student student
	//db.Model(&student).Where("name = ?", "张三").Update("grade", "四年级")

	//编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	var student []student
	db.Unscoped().Debug().Where("age < ?", 15).Delete(&student)
}
