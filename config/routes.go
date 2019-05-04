package config

import (
	"echt/handlers"

	"github.com/gin-gonic/gin"
)

// Route => Struct collects defined routes.
type Route struct {
	Method  string
	Route   string
	Handler gin.HandlerFunc
}

// Routes => Exported URL patterns.
var Routes = [...]Route{
	{"POST", "/submit", handlers.SubmitLink},
}
