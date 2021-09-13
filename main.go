package main

import (
	"net/http"

	"yuegefan/bd"

	"github.com/gin-gonic/gin"
)

var positions []string

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.POST("/result", func(c *gin.Context) {
		c.BindJSON(&positions)
		bdmap := bd.BDMap{}
		var locations []bd.Location
		for _, v := range positions {
			location, err := bdmap.Search(v)
			if err == nil {
				locations = append(locations, location)
			}
		}
		var center bd.Location
		for _, v := range locations {
			if center.Lat == 0 {
				center = v
			} else {
				center.Lat = (center.Lat + v.Lat) / 2
				center.Lng = (center.Lng + v.Lng) / 2
			}
		}
		results, err := bdmap.SearchCircle(center)
		if err != nil {
			c.JSON(http.StatusForbidden, err.Error())
			return
		}
		c.JSON(http.StatusOK, results)
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
