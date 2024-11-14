package main

import (
	"fmt"
	"gopi/api"
	"gopi/config"
	ds "gopi/datasources"
	"log"
	"os"
)

func main() {
	config := config.GetConfig()
	ds.InitDuckDB() // Create DuckDB
	os.Setenv("PORT", config.HttpPort)

	db := ds.InitDB(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DbUser, config.DbPassword,
		config.DbHost, config.DbPort, config.DbName))
	redis := ds.InitRedis(config.RedisPassword, config.RedisHost, config.RedisPort, 8)

	s3, err := ds.NewS3ClientWithCredentials(config.S3AcessKeyID, config.S3SecretKey, config.S3Region, config.S3Arn)

	if err != nil {
		log.Fatal(err)
	}

	duck := ds.GetDuckDB()

	api.Init(redis, db, s3, duck)
}
