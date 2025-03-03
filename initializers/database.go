package initializers

import (
	"fmt"
	"m-server-api/config"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	DB  *gorm.DB
	RDB *redis.Client
	Ctx = context.Background()
)

// 连接mysql
func connectMysql(cfg config.DatabaseConfig) {
	dsn := cfg.MySQL.User + ":" + cfg.MySQL.Password + "@tcp(" + cfg.MySQL.Host + ":" + fmt.Sprint(cfg.MySQL.Port) + ")/" + cfg.MySQL.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect to MySQL database")
	}

	db, err := DB.DB()
	if err != nil {
		panic(err)
	}

	// 设置连接池 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.SetMaxOpenConns(10)

	// 设置最大连接数 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	db.SetMaxIdleConns(60)

	// 设置最大连接超时
	db.SetConnMaxLifetime(time.Minute * 60)
}

// 连接redis
func connectRedis(cfg config.DatabaseConfig) {
	RDB = redis.NewClient(&redis.Options{
		Addr:         cfg.Redis.Host + ":" + fmt.Sprint(cfg.Redis.Port),
		Password:     cfg.Redis.Password,
		DB:           cfg.Redis.Db,
		MaxRetries:   3,
		DialTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 20,
		WriteTimeout: time.Second * 20,
		PoolSize:     50,
		MinIdleConns: 2,
		PoolTimeout:  time.Minute,
	})
	if err := RDB.Ping(Ctx).Err(); err != nil {
		panic(err)
	}
}

func CloseMysql() error {
	db, err := DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func CloseRedis() error {
	return RDB.Close()
}

func InitDatabase() {
	cfg := config.Get().Database

	connectMysql(cfg)
	connectRedis(cfg)
}
