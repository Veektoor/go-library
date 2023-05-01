package models

import(
	"github.com/jinzhu/gorm"
	 "go-library/pkg/config"
	 "gorm.io/driver/mysql"
	 "gorm.io/gorm"
)

var db *gorm.DB

type Book struct{
	gorm.model
	Name string `gorm:"json:" "name"`
	Author string `json: "author"`
	Publication string `json:"publication"`

}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	returb b
}

func GetAllBooks() []Books{
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64)(*Book, *gorm.DB){
	var getBook Book
	db:=db.Where ("ID=?", Id).Fand(&getBook)
	return &getBook.db
}
func DeleteBook(ID int64)Book{
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}