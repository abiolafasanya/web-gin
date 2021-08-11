package main

import "github.com/gin-gonic/gin"

func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":            "Hello world",
		"name of programmer": "Abiola Fasanya",
	})
}
