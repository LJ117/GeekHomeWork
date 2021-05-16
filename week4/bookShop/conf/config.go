package conf

import "fmt"

type ConfigInfo struct {
	ApiHost   string
	MySqlRoot string
}

type Role struct {
	BookStore int //书店管理者 对应：type 1
}

type RoleType struct {
	BookStore      int8 //书店管理者 对应：type 1
	BookStoreStaff int8 //书店下员工管理者 对应：type 11
}

var MySqlRoot string
var RoleInfo Role

func init() {

	RolePro := Role{
		BookStore: 10001,
	}

	configDev := ConfigInfo{
		ApiHost:   "http://127.0.0.1:7788/",
		MySqlRoot: "root:root@(127.0.0.1:3306)/book_shop?charset=utf8&parseTime=True&loc=Local",
	}

	fmt.Printf("开发环境：%+v\n", configDev)

	//切换环境配置
	configInfo := configDev

	fmt.Printf("当前连接环境：%+v\n", configInfo)
	fmt.Printf("当前固定角色配置环境：%+v\n", RoleInfo)

	MySqlRoot = configInfo.MySqlRoot
	RoleInfo = RolePro

}
