package admin

import (
	"bookShop.seven.com/admin/controller/bookController"
	"github.com/gin-gonic/gin"
)

func InitV1Router(router *gin.Engine) {

	//路由分组
	// Version 1
	v1 := router.Group("/api/v1/admin")
	{
		// 书店的CRUD
		bookShopRoute := v1.Group("/book")
		{
			bookShopRoute.GET("/find", bookController.GetInfo)
			//bookShopRoute.POST("/add", bookController.AddBook)
			//bookShopRoute.POST("/delete", bookController.DeleteBook)
			//bookShopRoute.POST("/update", bookController.UpdateBook)
		}

	}

}
