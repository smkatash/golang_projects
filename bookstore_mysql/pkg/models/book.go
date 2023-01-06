package models
import (
	"github.com/jinzhu/gorm"
	"github.com/smkatash/golang_projects/bookstore_mysql/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var b Book
	db := db.Where("ID=?", Id).Find(&b)
	return &b, db
}

func DeleteBook(Id int64) Book {
	var b Book
	db.Where("ID=?", Id).Delete(b)
	return b
}