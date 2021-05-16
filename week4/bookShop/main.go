package main

import (
	"bookShop.seven.com/conf"
	"bookShop.seven.com/database"
	"bookShop.seven.com/route/admin"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

func main() {
	//获取当前路径
	dir, err := os.Getwd()
	if err != nil {
		fmt.Print("获取当前路径失败，原因: ")
		panic(err)
	}
	conf.BasePath = dir

	dateT := time.Unix(time.Now().Unix(), 0).Format("2006-01-02")
	pathUrl := "./logs/" + dateT + "/"

	if err := os.MkdirAll(pathUrl, 0755); err != nil {
		return
	}

	//输出Gin日志到指定文件
	logfile, err := os.Create(pathUrl + "gin_http.log")
	if err != nil {
		fmt.Println("Could not create log file")
	}
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.MultiWriter(logfile)

	//使用默认中间件创建一个gin路由器
	router := gin.Default()

	//初始化后台管理admin 路由 v1版本
	admin.InitV1Router(router)

	defer func() {
		_ = database.Conn.Close()
	}()

	//获取数据库配置的启动端口号
	port := "9527"

	fmt.Println(fmt.Sprintf("监听端口号是:%s", port))
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		fmt.Print(fmt.Sprintf("监听%s端口失败，失败原因：", port))
		panic(err)
	}
}
