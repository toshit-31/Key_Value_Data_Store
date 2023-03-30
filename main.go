package main

import (
	"in-memory-db/1/components"
	"in-memory-db/1/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QueryString struct {
	Query string `json:"query"`
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func valueResponse(res components.Result) gin.H {
	return gin.H{"value": res.Value}
}

var dataStore = new(core.DataStore)

func init() {
	dataStore.Init()
}

func main() {
	router := gin.Default()

	router.POST("/query", func(c *gin.Context) {
		var data QueryString
		err := c.BindJSON(&data)
		if err != nil {
			c.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		result := dataStore.Query(data.Query)
		if result.Err != nil {
			println(result.Err.Error())
			c.JSON(http.StatusBadRequest, errorResponse(result.Err))
		} else {
			c.JSON(http.StatusOK, valueResponse(result))
		}
	})
}
