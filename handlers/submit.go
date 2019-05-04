package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SubmitLink(c *gin.Context) {

	c.JSON(http.StatusCreated, gin.H{"ok": true})
}
