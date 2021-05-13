package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var urlMap = make(map[string]string)

func main() {
	router := gin.Default()

	// opt := badger.DefaultOptions("").WithInMemory(true)
	// db, err := badger.Open(opt)
	// if err != nil {
	//     log.Fatal(err)
	// }
	// defer db.Close()

	urlMap["google"] = "https://google.com"
	router.GET("/:urlname", redirectGet)
	router.POST("/:urlname", redirectSet)
	router.GET("/list", listURLs)

	router.Run(":8080")
}

func redirectGet(c *gin.Context) {
	// Look up
	url := urlMap[c.Param("urlname")]
	c.Redirect(http.StatusMovedPermanently, url)
}

func redirectSet(c *gin.Context) {
	// Set URL
	redirectURL := c.PostForm("redirectURL")
	urlMap[c.Param("urlname")] = redirectURL
}

func listURLs(c *gin.Context) {
	var message string
	for urlName, url := range urlMap {
		message += fmt.Sprintf("%s: %s\n", urlName, url)
	}
	c.String(http.StatusOK, "%s", message)
}