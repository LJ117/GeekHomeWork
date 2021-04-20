package main

import (
	"database/sql"
	"github.com/pkg/errors"
)

//  1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
// Answer: 应该Wrap, 因为 dao 层的应用会被 service 层等调用, dao层并非程序的顶部. 它被调用后会有其他库协作的行为.

func example() (err error) {
	open, openErr := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if openErr != nil {
		return errors.Wrap(err, "mysql: openDB error")
	}
	q := "select   * from    test where test.id = ? ; "
	_, readErr := open.Query(q, -1)
	if errors.Is(readErr, sql.ErrNoRows) {
		return errors.Wrap(err, "mysql: Not found this Row")
	}
	return
}
