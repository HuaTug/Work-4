package configs

import (
	"4-TikTok/cache"
	"4-TikTok/model"
	"context"
	"fmt"
	"github.com/go-redis/redis"
	logging "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/ini.v1"
	"os"
	"strconv"
	"strings"
)

var (
	Db         string
	DbPort     string
	DbHost     string
	DbUser     string
	DbPassword string
	DbName     string
)

var (
	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

var (
	OssName     string
	OssPassword string
)

var (
	MongoDBClient *mongo.Client
	MongoDBName   string
	MongoDBAddr   string
	MongoDBPwd    string
	MongoDBPort   string
)

func Init() {
	file, err := ini.Load("./configs/conf.ini")
	if err != nil {
		executablePath, _ := os.Executable()
		fmt.Println("这是路径" + executablePath)
		panic(err)
	}
	LoadMysqlData(file)
	LoadRedis(file)
	LoadMinio(file)
	LoadMongoDB(file)
	path := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	model.Database(path)
	//Redis()
	MongoDB()
}
func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadRedis(file *ini.File) {
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}

func LoadMinio(file *ini.File) {
	OssName = file.Section("minio").Key("OssName").String()
	OssPassword = file.Section("minio").Key("OssPassword").String()
}

func Redis() {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPw,
		DB:       int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		logging.Info(err)
		panic(err)
	}
	cache.RedisClient = client
	fmt.Println("Redis连接成功")
}

func MongoDB() {
	// 设置mongoDB客户端连接信息
	clientOptions := options.Client().ApplyURI("mongodb://" + MongoDBAddr + ":" + MongoDBPort)
	var err error
	MongoDBClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logging.Info(err)
	}
	err = MongoDBClient.Ping(context.TODO(), nil)
	if err != nil {
		logging.Info(err)
	}
	logging.Info("MongoDB Connect")
}

func LoadMongoDB(file *ini.File) {
	MongoDBName = file.Section("MongoDB").Key("MongoDBName").String()
	MongoDBAddr = file.Section("MongoDB").Key("MongoDBAddr").String()
	MongoDBPwd = file.Section("MongoDB").Key("MongoDBPwd").String()
	MongoDBPort = file.Section("MongoDB").Key("MongoDBPort").String()
}
