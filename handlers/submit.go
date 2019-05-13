package handlers

import (
	"echt/core"
	"echt/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type submitLinkRequest struct {
	Link string `json:"link"`
}

// SubmitLink => Receive POSTed link data.
func SubmitLink(c *gin.Context) {
	var request submitLinkRequest

	err := c.ShouldBind(&request)
	if nil != err {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request."})
	}

	var link = request.Link
	record := models.Content{Link: link}

	core.GetDbConn().Create(&record)

	c.JSON(http.StatusCreated, gin.H{"ok": true})
}


func ListSubmitted(c *gin.Context) {
	var submitted []models.Content

	core.GetDbConn().Find(&submitted)

	c.JSON(http.StatusOK, submitted)
}
