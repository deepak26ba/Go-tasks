package main

import (
	"brevo/api/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("check")

	g := gin.Default()

	routes.Routes(g)

	g.Run(":8080")
}
