package api

import (
	"jkt48scrap/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func member(c *gin.Context) {
	data := database.GetAllData()

	c.JSON(200, gin.H{"member": data})

}

func memberID(c *gin.Context) {

	data := database.GetName(c.Param("nama"))

	c.JSON(200, gin.H{"member": data})

}

func myRoute(r *gin.RouterGroup) {
	r.GET("/member", member)
}

func init() {
	app = gin.New()
	r := app.Group("/api-v1")
	myRoute(r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
