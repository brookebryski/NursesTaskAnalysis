package routePack

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routing() {
	router := gin.Default()
	router.LoadHTMLGlob("template/*")
	// http.Handle("/login/", http.FileServer(http.Dir("./nurseLogin")))
	// fmt.Println("Hello World")
	router.GET("/login/Nurses", func(c *gin.Context) {
		c.HTML(http.StatusOK, "nurseLogin.html", gin.H{
			"content": "This is a login page...",
		})
	})
	router.GET("/login/Supervisors", func(c *gin.Context) {
		c.HTML(http.StatusOK, "supLogin.html", gin.H{
			"content": "This is a login page...",
		})
	})
	router.GET("/tasks/input", func(c *gin.Context) {
		c.HTML(http.StatusOK, "nurseTasks.html", gin.H{
			"content": "This the Nurses Task Page...",
		})
	})
	router.GET("/review/tasks", func(c *gin.Context) {
		c.HTML(http.StatusOK, "supTasksView.html", gin.H{
			"content": "This is the supervisors View Page...",
		})
	})
	router.Run("localhost:3000")
}
