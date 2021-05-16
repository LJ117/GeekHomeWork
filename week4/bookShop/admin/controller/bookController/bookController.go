package bookController

import (
	"bookShop.seven.com/admin/model/bookModel"
	"bookShop.seven.com/common"
	"github.com/gin-gonic/gin"
)

type ListParam struct {
	BookID int `json:"book_id" validate:"required"`
}

func GetInfo(c *gin.Context) {
	var param ListParam
	err := c.ShouldBindQuery(&param)
	if err != nil {
		common.Error(c, -1, err.Error(), []string{})
		return
	}

	var model bookModel.Book
	model.BID = param.BookID

	info, err := model.GetBookInfo()
	if err != nil {
		common.Error(c, -1, err.Error(), []string{})
		return
	}

	common.Success(c, 0, "success", info)
}
