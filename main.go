package main

import (
	"github.com/gin-gonic/gin"
	"github.com/murcagr/ReviewApi/routes"
)

func main() {

	r := gin.Default()

	r = routes.SetupRouter(r)

	r.Run()

}
