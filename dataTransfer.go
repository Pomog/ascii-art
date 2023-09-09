package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type message struct {
	ID        int    `json:"id"`
	Body      string `json:"body"`
	Completed bool   `json:"completed"`
}

type ViewData struct {
	Strings []string
}

var messages = []message{
	{ID: 1, Body: "Hello World", Completed: false},
	{ID: 2, Body: "Sleep", Completed: false},
	{ID: 3, Body: "Eat", Completed: false},
}

func RenderStringsPage(c *gin.Context) {
	// asciiartSlice := []string{"Art1", "Art2", "Art3"}
	asciiartSlice2 := AsciiArt
	// Create a ViewData struct with the strings data
	data := ViewData{
		Strings: asciiartSlice2,
	}

	// Render the HTML template with the data
	c.HTML(http.StatusOK, "asciiart.html", data)
}

func RenderTestPage(c *gin.Context) {
	fmt.Println("RenderTestPage called") // Add this line for debugging
	// Render the HTML template from the file
	c.HTML(http.StatusOK, "test.html", nil)
}

func GetMessages(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, messages)
}

func AddMessage(context *gin.Context) {
	var newMessage message

	if err := context.BindJSON(&newMessage); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	messages = append(messages, newMessage)
	fmt.Printf("messages: %v\n", messages)
	context.IndentedJSON(http.StatusCreated, newMessage)
}

func RenderTestPage2(c *gin.Context) {
	c.HTML(http.StatusOK, "test.html", nil)
}
