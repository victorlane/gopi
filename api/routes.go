package api

import (
	"math/rand"
	"net/http"
	"time"

	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var multiplier = 8

func Init(redis *redis.Client) {
	router := gin.Default()
	redisStore := persist.NewRedisStore(redis)

	{
		v1 := router.Group("/v1")
		v1.GET("/ping", cache.CacheByRequestURI(redisStore, 2*time.Minute), func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": multiplier * rand.Intn(multiplier*2),
			})
		})
	}

	router.Run()
}
