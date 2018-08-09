package main

import (
	"openRPA-log-server/app/models"
	"openRPA-log-server/app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.InitDB()
	routes.InitRoutes(r)

	r.Run(":9000") // listen and serve on 0.0.0.0:9000
}
