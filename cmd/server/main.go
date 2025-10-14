package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lapanxd/volatus-api/internal/routes"
)

func main() {
	r := gin.Default()

	routes.HealthRoutes(r)

	r.Run(":8080")

}
