package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// create table books
// (
// id         bigint auto_increment
// primary key,
// title       varchar(25) null,
// author varchar(25) null,
// price     double      null
// );
type Book struct {
	ID     int8    `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float32 `db:"price"`
}

func RunSqlX2() {
	db, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gorm")
	if err != nil {
		fmt.Println("连接出错: ", err)
	}

	//插入数据
	//b1 := Book{Title: "Java语言开发", Author: "谭浩强", Price: 28.5}
	//b2 := Book{Title: "数据结构", Author: "严佬", Price: 50}
	//b3 := Book{Title: "控制工程论", Author: "钱佬", Price: 68.5}

	//books := []*Book{&b1, &b2, &b3}
	//
	//sqlStr := "insert into books (title, author, price) values(:title, :author, :price)"
	//_, err = db.NamedExec(sqlStr, books)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
	sqlStr := "select * from  books where price > 50"
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println("查询失败", err)
	}
	defer rows.Close()
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(book)
	}

}
