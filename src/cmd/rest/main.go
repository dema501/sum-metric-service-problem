package main

import (
	"net/http"
	"sum-metric-service-problem/src/internal/memstorage"
	"sum-metric-service-problem/src/internal/rest"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	store := memstorage.NewStorage()
	defer store.Purge()

	rs := rest.Server{
		ServerPort: "5000",
	}

	rs.Run(
		func(rg *gin.RouterGroup) {
			rg.POST("/metric/:name", PostMetricHandler(store))
			rg.GET("/metric/:name/sum", GeSumMetricHandler(store))
		},
	)
}

// GeSumMetricHandler hande GET /metric/{key}/sum
func GeSumMetricHandler(store *memstorage.Storage) func(c *gin.Context) {
	return func(c *gin.Context) {
		backet := c.Params.ByName("name")

		sum, err := store.GetSumItemsAddedBefore(backet, 1*time.Hour)
		if err == memstorage.ErrNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"value": sum,
		})
	}
}

// PostMetricHandler hande POST  /metric/{key}
func PostMetricHandler(store *memstorage.Storage) func(c *gin.Context) {
	return func(c *gin.Context) {
		backet := c.Params.ByName("name")
		var metric rest.AddMetricRequest

		if err := c.BindJSON(&metric); err != nil {
			c.AbortWithStatusJSON(400, gin.H{
				"message": err,
			})
			return
		}

		store.AddToBacket(backet, metric.Value)

		c.JSON(http.StatusOK, gin.H{})
	}
}
