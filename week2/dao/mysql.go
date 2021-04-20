package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	_ "gorm.io/driver/mysql"
)

//  1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
// Answer: 应该Wrap, 因为 dao 层的应用会被 service 层等调用, dao层并非程序的顶部. 它被调用后会有其他库协作的行为.

func main() {
	//fmt.Printf("%+v",example())
	err := example()
	if err != nil {
		//fmt.Printf("%v\n",err)
		fmt.Printf("%+v\n",err)
	}
}

func example() (err error) {
	open, openErr := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/geek?charset=utf8mb4&parseTime=True&loc=Local")
	//open, openErr := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/geek?charset=utf8mb4&parseTime=True&loc=Local")
	if openErr != nil {
		err = errors.New(" openDB error")
	}
	q := "select   * from    test where test.id = -1 ; "
	row := open.QueryRow(q)
	type Temp struct {
		id   int
		name string
	}
	var t Temp
	tErr := row.Scan(&t)
	if errors.Is(tErr, sql.ErrNoRows) {
		err = errors.New(" Not found this Row")
	}

	return errors.Wrap(err, "dao: ")
}
