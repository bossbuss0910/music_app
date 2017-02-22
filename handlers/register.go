package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"music_app/models"
	"net/http"
)

type registerHandlerInterface interface {
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
}

func NewRegisterHandler() registerHandlerInterface {
	return &registerHandler{}
}

type registerHandler struct {
}

func (h *registerHandler) RegisterUser(c *gin.Context) {
	db, err := gorm.Open("mysql", "root@/music_app?charset=utf8&parseTime=True")
	if err != nil {
	  panic("failed to connect database")
	}
	user := new(models.User)
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	db.Create(user)
	c.JSON(http.StatusCreated, gin.H{"status": "ok"})
}

func (h *registerHandler) LoginUser(c *gin.Context) {
	db, err := gorm.Open("mysql", "root@/music_app?charset=utf8&parseTime=True")
	if err != nil {
	  panic("failed to connect database")
	}
	user := new(models.User)
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var users models.User
	db.Where(user).First(&users);
	if users.Id == 0 {
		c.JSON(http.StatusCreated, gin.H{"status": "bad"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success"})
}
