package dao

import (
	"database/sql"
)

//  1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

func init() {
	open, err2 := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/geek?charset=utf8mb4&parseTime=True&loc=Local")

}
