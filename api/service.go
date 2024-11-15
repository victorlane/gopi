package api

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	ds "gopi/datasources"

	jwt "github.com/appleboy/gin-jwt/v2"
	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var multiplier = 8
var cacheOpts = cache.WithPrefixKey("gopi")

func Init(redis *redis.Client, mysql *sql.DB, s3 *ds.S3Client, duck *sql.DB, jwtSecret *string) {
	// Set production mode
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	authMiddleware, err := jwt.New(initAuthParams(*jwtSecret))
	fmt.Println(authMiddleware)
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	router.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		_, err := duck.Exec(`
			INSERT INTO logs (client_ip, timestamp, method, path, protocol, status_code, latency, user_agent, created_at)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
		`, c.ClientIP(), time.Now(), c.Request.Method, c.Request.URL.Path, c.Request.Proto, status, latency.String(), c.Request.UserAgent(), time.Now())
		if err != nil {
			fmt.Printf("Error inserting log into DuckDB: %v\n", err)
		}
	})

	redisStore := persist.NewRedisStore(redis)

	api := router.Group("/v1")
	{
		api.GET("/ping", cache.CacheByRequestURI(redisStore, 2*time.Minute, cacheOpts), func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": multiplier * rand.Intn(multiplier*2),
			})
		})
	}

	router.NoRoute(gin.WrapH(http.FileServer(http.Dir("public"))))

	router.Run()
}
