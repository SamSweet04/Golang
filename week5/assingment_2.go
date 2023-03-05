package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type System interface {
	Register(Registration)
	SignIn(Database)
	AddItem(string,float64)
	SearchItem(string)
	FilteringItems(float64,float64,float64,float64)
	Rate(Database, Authorization, string, float64)

}

type Registration struct {
	Name     string
	Surname  string
	Age      int
	Login    string
	Password string
}

type Authorization struct {
	Login    string
	Password string
}

type Item struct {
	Name       string
	Price      float64
	Rating     float64
	RatingList []float64
}

type Database struct {
	Logins []Registration
	Items  []Item
}

func (d *Database) RegisterHandler(c *gin.Context) {
	var r Registration
	err := c.ShouldBindJSON(&r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reg := d.Register(r)
	if reg == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "user registered"})
	}
}

func (a *Authorization) SignInHandler(c *gin.Context) {
	var d Database
	err := c.ShouldBindJSON(&d)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	msg := a.SignIn(d)
	if msg == "No authorized!!!" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid login credentials"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "user signed in"})
	}
}

func (d *Database) AddItemHandler(c *gin.Context) {
	var i Item
	err := c.ShouldBindJSON(&i)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item := d.AddItem(i.Name, i.Price)
	c.JSON(http.StatusOK, gin.H{"message": "item added", "item": item})
}

func (d *Database) SearchItemHandler(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing required parameter 'name'"})
		return
	}
	item := d.SearchItem(name)
	if item == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"item": item})
	}
}

func (d *Database) FilteringItemsHandler(c *gin.Context) {
	price1Str := c.Query("price1")
	price2Str := c.Query("price2")
	rating1Str := c.Query("rating1")
	rating2Str := c.Query("rating2")
	price1, err := strconv.ParseFloat(price1Str, 64)
	if err != nil {
		price1 = 0
	}
	price2, err := strconv.ParseFloat(price2Str, 64)
	if err != nil {
		price2 = 0
	}
	rating1, err := strconv.ParseFloat(rating1Str, 64)
	if err != nil {
		rating1 = 0
	}
	rating2, err := strconv.ParseFloat(rating2Str, 64)
	if err != nil {
		rating2 = 0
	}
	items :=
