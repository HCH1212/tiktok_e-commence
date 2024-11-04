package dao

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB
var RDB *redis.Client

// 用gorm初始化mysql
func InitMysql() {
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	database := viper.GetString("datasource.database")
	charset := viper.GetString("datasource.charset")
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	DB = db
	//err = DB.AutoMigrate(&model.User{}) //自动生成表
	//if err != nil {
	//	log.Println(err)
	//}
}

func InitRedis() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis服务器地址和端口
		Password: "",               // 如果没有密码则为空字符串
		DB:       0,                // 使用的数据库编号，默认为0
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Println("Failed to connect to Redis:", err)
		return
	}
	RDB = rdb
}
