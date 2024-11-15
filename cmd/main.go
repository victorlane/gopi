package main

import (
	"gopi/api"
	"gopi/config"
	ds "gopi/datasources"
	"log"
	"os"
)

func main() {
	config := config.GetConfig()
	os.Setenv("PORT", config.HttpPort)

	ds.InitDuckDB(config.Db)
	db := ds.InitDB(config.Db)
	redis := ds.InitRedis(config.RedisPassword, config.RedisHost, config.RedisPort, 8)
	s3, err := ds.NewS3ClientWithCredentials(config.S3AcessKeyID, config.S3SecretKey, config.S3Region, config.S3Arn)

	if err != nil {
		log.Fatal(err)
	}

	duck := ds.GetDuckDB()

	RunCrons(duck, db)

	api.Init(redis, db, s3, duck, &config.JWTSecretKey)
}
