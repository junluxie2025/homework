package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//create table employees
//(
//    id         bigint auto_increment
//        primary key,
//    name       varchar(25) null,
//    department varchar(25) null,
//    salary     double      null
//);

type Employee struct {
	ID         string  `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float32 `db:"salary"`
}

func RunSqlX1() {
	db, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gorm")
	if err != nil {
		fmt.Println("连接出错: ", err)
	}

	//e1 := Employee{Name: "张三", Department: "技术部", Salary: 1000}
	//e2 := Employee{Name: "李四", Department: "技术部", Salary: 2000}
	//e3 := Employee{Name: "王五", Department: "市场", Salary: 3000}
	//employees := []*Employee{&e1, &e2, &e3}
	//插入数据
	//sqlStr := "insert into employees (name, department, salary) values(:name, :department, :salary)"
	//_, err = db.NamedExec(sqlStr, employees)
	//defer db.Close()
	//if err != nil {
	//	fmt.Println(err)
	//}

	//编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，
	//并将结果映射到一个自定义的 Employee 结构体切片中。
	//sqlStr := "select * from employees where department = ?"
	//rows, err := db.Query(sqlStr, "技术部")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer rows.Close()
	//var emps []Employee
	//for rows.Next() {
	//	var e Employee
	//	err = rows.Scan(&e.ID, &e.Name, &e.Department, &e.Salary)
	//	if err != nil {
	//		fmt.Println("遍历出错: ", err)
	//		return
	//	}
	//	emps = append(emps, e)
	//}
	//for _, emp := range emps {
	//	fmt.Println(emp)
	//}

	//编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
	sqlStr := "select * from  employees order by salary desc limit 1"
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var e Employee
		err = rows.Scan(&e.ID, &e.Name, &e.Department, &e.Salary)
		if err != nil {
			fmt.Println("遍历出错: ", err)
			return
		}
		fmt.Println(e)
	}

}
