package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	//SQL语句练习 题目1：基本CRUD操作
	//RunCRUD(db)

	//SQL语句练习 题目2:事务语句
	//RunTransaction(db)

	//Sqlx入门 题目1：使用SQL扩展库进行查询
	//RunSqlX1()

	//Sqlx入门 题目1：实现类型安全映射
	//RunSqlX2()

	//进阶gorm 题目1：模型定义
	RunGrom1(db)

	//进阶gorm   题目2：关联查询
	RunGrom2(db)

	//进阶gorm   题目3：钩子函数
	RunGrom3(db)
}
