package handlers

import (
	"github.com/gin-gonic/gin"
)

// Route => Collects defined routes.
type Route struct {
	Method  string
	Route   string
	Handler gin.HandlerFunc
}

// Routes => Exported URL patterns.
var Routes = [...]Route{
	{"POST", "/submit", SubmitLink},
	{"GET", "/submitted", ListSubmitted},
}
