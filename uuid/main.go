package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

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
	if err != nil {
		c.Error(err)
	}
	book := Book{}
	book.BeforeCreate()
	if c.Request.ContentLength != 0 {
		err = json.Unmarshal(body, &book)
		if err != nil {
			c.Error(err)
		}
		userCreated := book.SaveBook()
		if err != nil {
			c.Error(err)
		}
		c.JSON(200, userCreated)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pass a body"})
	}
}

//GetParams ...
func GetParams(c *gin.Context) {

	vars := c.Param("uuid")
	spars := c.Param("lol")
	userid, err := uuid.Parse(vars)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse UUID"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"uuid": userid, "lol": spars})
}

func main() {
	router := gin.Default()
	router.POST("/", CreateBook)
	router.GET("/getter/:uuid/:lol", GetParams)
	router.Run(":8000")
}
