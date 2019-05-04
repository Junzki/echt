package handlers

import (
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
	c.JSON(http.StatusCreated, gin.H{"link": link})
}
