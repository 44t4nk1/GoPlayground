package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//Book ...
type Book struct {
	UUID        uuid.UUID `gorm:"primary_key" json:"uuid"`
	Title       string    `gorm:"size:255;not null;unique" json:"title"`
	Author      string    `gorm:"size:255;not null;unique" json:"author"`
	ImageURL    string    `gorm:"size:255;not null;unique" json:"image_url"`
	Description string    `gorm:"size:255;not null;unique" json:"description"`
}

//BeforeCreate ...
func (b *Book) BeforeCreate() {
	newuuid := uuid.New()
	b.UUID = newuuid
}

//SaveBook ...
func (b *Book) SaveBook() *Book {
	return b
}

//CreateBook ...
func CreateBook(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	book := Book{}
	book.BeforeCreate()
	err = json.Unmarshal(body, &book)
	if err != nil {
		c.Error(err)
	}
	userCreated := book.SaveBook()
	if err != nil {
		c.Error(err)
	}
	c.JSON(200, userCreated)
}

func main() {
	router := gin.Default()
	router.POST("/", CreateBook)
	router.Run(":8000")
}
