package controllers

import (
	"brainstorm-press-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GET /posts
func FindPosts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var posts []models.Post
	db.Find(&posts)

	c.IndentedJSON(http.StatusOK, gin.H{"data": posts})
}

// POST /posts
func CreatePost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input models.CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create post
	post := models.Post{Title: input.Title, Author: input.Author, Body: input.Body}
	db.Create(&post)

	c.IndentedJSON(http.StatusOK, gin.H{"data": post})
}

// GET /posts/:id
func FindPost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var post models.Post
	if err := db.First(&post, c.Param("id")).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": post})
}

// PATCH /posts/:id
func UpdatePost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Find post
	var post models.Post
	if err := db.First(&post, c.Param("id")).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// Validate input
	var input models.UpdatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&post).Updates(input)

	c.IndentedJSON(http.StatusOK, gin.H{"data": post})
}

// DELETE /posts/:id
func DeletePost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var post models.Post
	if err := db.First(&post, c.Param("id")).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	db.Delete(&post)

	c.IndentedJSON(http.StatusOK, gin.H{"data": true})
}
