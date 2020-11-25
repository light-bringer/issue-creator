package controllers

import (
	"net/http"

	"issue-creator/db"
	"issue-creator/models"

	"github.com/gin-gonic/gin"
)

func FindIssues(c *gin.Context) {
	var issues []models.Issue
	db.DB.Find(&issues)

	c.JSON(http.StatusOK, gin.H{"data": issues})
}
