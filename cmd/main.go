package main

import (
	"fmt"
	"gopi/api"
	"gopi/config"
	ds "gopi/datasources"
	"log"
)

func main() {
	config := config.GetConfig()
	db := ds.InitDB(fmt.Sprintf("mysql://%s:%s@%s:%s/%s", config.DbUser, config.DbPassword,
		config.DbHost, config.DbPort, config.DbName))
	redis := ds.InitRedis(config.RedisPassword, config.RedisHost, config.RedisPort, 8)

	s3, err := ds.NewS3ClientWithCredentials(config.S3AcessKeyID, config.S3SecretKey, config.S3Region, config.S3Arn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(s3)
	fmt.Print(db)

	api.Init(redis)
}
