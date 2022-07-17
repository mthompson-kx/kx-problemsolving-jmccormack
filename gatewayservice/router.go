package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"encoding/json"
)

var num = 0
const servers = 3

func selectServer()int{
	
	var tmp = num % servers
	num+=1
	return tmp
}

func dataRouting(root *gin.Engine) {

	server := root.Group("")
	server.GET("/data", data)

}

func data(c *gin.Context) {

	fmt.Println("Retrieving data.json...")

	var res map[string]interface{}

	retry := servers

	for retry > 0 {
		fmt.Printf("Attempt %d\n", 3-retry+1)
		retry-=1

		serverNo := selectServer()
		fmt.Printf("Selecting from server %d\n", serverNo+1)

		address := fmt.Sprint("http://storageservice", serverNo+1, ":8000/data.json")
		resp, err := http.Get(address)

		if err != nil {
			continue
		}

		defer resp.Body.Close()

		json.NewDecoder(resp.Body).Decode(&res)
	
		c.JSON(http.StatusOK, res)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Error": "Internal Server Error",
	})
	return

}

func statusRouting(root *gin.Engine) {	
	
	server := root.Group("")
	server.GET("/status", status)
}

func status(c *gin.Context) {

	fmt.Println("Status check...")

	statuses := make(map[string]string)

	for i := 1; i < 4; i++ {

		address := fmt.Sprint("http://storageservice", i, ":8000/health")
		_, err := http.Get(address)

		if err != nil {			
			statuses[fmt.Sprint("Server ", i)] = "unhealthy"
			continue
		}

		statuses[fmt.Sprint("Server ", i)] = "healthy"

	}

	c.JSON(http.StatusOK, statuses)
}

func routerEngine() *gin.Engine {
	
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	dataRouting(router)
	statusRouting(router)
	router.NoRoute(
		func(c *gin.Context) {
			if c.Request.Method == http.MethodOptions {
				c.String(http.StatusOK, "")
				return
			}
			c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
		},
	)
	return router
}
