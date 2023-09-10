package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewData struct {
	Strings []string
}

func RenderStringsPage(c *gin.Context) {
	asciiartSlice2 := AsciiArt
	// Create a ViewData struct with the strings data
	data := ViewData{
		Strings: asciiartSlice2,
	}

	// Render the HTML template with the data
	c.HTML(http.StatusOK, "asciiart.html", data)
}
