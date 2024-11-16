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
	db := ds.GetDB(config.Db)
	duck := ds.GetDuckDB()

	if !ds.IsInitialized(duck) {
		ds.InitDB(db)
		ds.InitDuckDB(config.Db, duck)
	}

	redis := ds.InitRedis(config.RedisPassword, config.RedisHost, config.RedisPort, 8)
	s3, err := ds.NewS3ClientWithCredentials(config.S3AcessKeyID, config.S3SecretKey, config.S3Region, config.S3Arn)
	if err != nil {
		log.Fatal(err)
	}

	RunCrons(duck, db)
	api.Init(redis, db, s3, duck, &config.JWTSecretKey)
}
