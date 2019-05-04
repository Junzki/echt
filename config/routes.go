package config

import (
	"echt/handlers"
	"github.com/gin-gonic/gin"
)

// Route collects defined routes.
type Route struct {
	Method  string
	Route   string
	Handler gin.HandlerFunc
}

var Routes = [...]Route{
	{"GET", "/submit", handlers.SubmitLink},
}

