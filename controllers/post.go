package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sid-008/Postgres_CRUD/database"
	"github.com/sid-008/Postgres_CRUD/helper"
	"github.com/sid-008/Postgres_CRUD/models"
)

func AddPost(c *gin.Context) {
	var input models.Post
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.UserID = user.ID

	savedEntry, err := input.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}

func GetAllPosts(c *gin.Context) {
	user, err := helper.CurrentUser(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user.Posts})
}

func UpdateOnePost(c *gin.Context) {
	var update models.Post
	if err := c.ShouldBindJSON(&update); err != nil { // bind update request
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(c) //auth to get the current user

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post := user.Posts

	database.Database.Model(&post).Where("Title = ?", update.Title).Update("Content", update.Content)
}

func GetAllPostsAnon(c *gin.Context) {
	posts := models.GetAll()
	log.Println(posts) //TODO add error handling
	c.JSON(http.StatusOK, gin.H{"data": posts})
}
