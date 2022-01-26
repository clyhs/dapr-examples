package model

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"dapr-examples/hello-service/common/config"
)

var (
	DB *gorm.DB
)

func InitDB() {
	fmt.Println("init entrance model...")

	// 从配置文件中获取连接信息
	host := config.String("database.host")
	port := config.String("database.port")
	user := config.String("database.user")
	password := config.String("database.password")
	name := config.String("database.name")
	charset := config.String("database.charset")
	prefix := config.String("database.prefix")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local", user, password, host, port, name, charset)
	//dsn := "root:root@tcp(127.0.0.1:3306)/mps?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: prefix, // 表名前缀，`User` 的表名应该是 `mps_users`
			//SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `mps_user`
		},
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	maxIdle := config.Int("database.max-idle")         // 连接池最大闲置的连接数
	maxOpen := config.Int("database.max-open")         // 连接池最大打开的连接数
	maxLifetime := config.Int("database.max-lifetime") // 连接对象可重复使用的时间长度(秒)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(maxIdle)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(maxOpen)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(maxLifetime))

	autoMigrate(db)

	if config.Bool("database.debug") {
		db = db.Debug()
	}
	DB = db
	// defer close
	//i, err := db.DB()
	//defer i.Close()
}

func autoMigrate(db *gorm.DB) {
	// migrate model
	err := db.AutoMigrate(
		User{},     // user
	)
	if err != nil {
		//global.GVA_LOG.Error("register table failed", zap.Any("err", err))
		log.Fatalf("migrate error: %v", err.Error())
		os.Exit(0)
	}
}
