package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

var podIP = os.Getenv("POD_OP")

func main() {
	r := gin.Default()

	r.GET("/",
		func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "bye",
				"ip":      podIP,
			})
		},
	)

	r.Run(":3001")
}
