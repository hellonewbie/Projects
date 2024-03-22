package initMatters

import (
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"smallProGarm/conn"
	"smallProGarm/utils"
)

var (
	DB     *gorm.DB
	Client *redis.Client
)

func InitMatters() (*gorm.DB, *redis.Client) {
	//初始化之后还返回了一个(MySql)数据库连接的对象DB
	db := new(conn.Sp)
	dbUser := db.SetDbUserName("root")
	dbPwd := db.SetDbPassword("ljy12345...")
	dbPort := db.SetDbPort("3306")
	dbHost := db.SetDbHost("127.0.0.1")
	dbdb := db.SetDbDataBase("program")
	DBInit, err := conn.InitMysql(dbUser, dbPwd, dbPort, dbHost, dbdb)
	DB = DBInit
	if err != nil {
		log.Print("InitMysql failed")
		log.Fatal(err)
	}
	//Redis初始化
	Redis := new(conn.RedisClient)
	addr := Redis.SetRedisAddr("175.178.72.197:6379")
	pwd := Redis.SetRedisPwd("ljy123...")
	Reidsdb := Redis.SetRedisDb(0)
	Client, _ = Redis.RedisInit(addr, pwd, Reidsdb)
	//第三个参数是存在时常
	//Client.Set("d", 1234, time.Minute)
	//MyLog初始化
	MyLogParam := new(utils.MLogParam)
	MyLogParam.MLogInit()
	//token初始化
	return DB, Client
}
