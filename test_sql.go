package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() error {
	dsn := "root:root@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
	db, _ = sql.Open("mysql", dsn)
	err := db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func insert() {
	sqlstr := "insert into user_tbl(username,password) values (?,?)"

	_, err := db.Exec(sqlstr, "zhangsan", "123456")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Println("insert success")
	}
}

type User struct {
	id       int
	name     string
	password string
}

func queryOneRow() {
	s := "select * from user_tbl"
	var user User
	err := db.QueryRow(s).Scan(&user.id, &user.name, &user.password)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("user: %v\n", user)
	}

}

func Exec() {
	s := "update user_tbl set username = ?,password = ? where id = ?"
	db.Exec(s, "aaaaa", "bbbbb", 1)
}
func queryAll() {
	s := "select * from user_tbl"
	var user User
	r, _ := db.Query(s)
	for r.Next() {
		r.Scan(&user.id, &user.name, &user.password)
		fmt.Printf("user: %v\n", user)
	}
}
func main() {
	// d, err := sql.Open("mysql", "root:root@/go_db")
	// if err != nil {
	// 	fmt.Println("open db error")
	// }
	// d.SetConnMaxLifetime(time.Minute * 3)
	// d.SetMaxOpenConns(10)
	// d.SetMaxIdleConns(10)
	err := InitDB()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Println("success")
	}
	fmt.Printf("db: %v\n", db)
	Exec()
	queryAll()
}
