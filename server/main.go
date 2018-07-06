package main

import (
	"net/http"

	"github.com/jeroenouw/AngularGolang/server/data"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/users", data.Index)
	http.HandleFunc("/users/show", data.Show)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

// func main() {
// 	router := gin.Default()
// 	router.LoadHTMLGlob("templates/*")
// 	router.GET("/", func(c *gin.Context) {
// 		c.HTML(http.StatusOK, "index.gohtml", gin.H{
// 			"title": "Main website",
// 		})
// 	})
// 	router.Run(":8080")
// }
