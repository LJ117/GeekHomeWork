package bookModel

import "bookShop.seven.com/database"

type Book struct {
	BID    int
	Name   string
	Author string
	Year   string
	Price  float32
}

func (b *Book) GetBookInfo() (Book, error) {
	var res Book

	db := database.Conn.Model(res)
	if err := db.Where("id = ?", b.BID).First(&res).Error; err != nil {
		return Book{}, err
	}

	return res, nil
}

func (b Book) TableName() string {
	return "book"
}
