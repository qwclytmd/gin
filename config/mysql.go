package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionMySQL(dbname string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", v.GetString("database.username"), v.GetString("database.password"), v.GetString("database.host"), v.GetInt("database.port"), dbname)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn + "?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                            // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                           // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                           // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                           // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                          // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	if err != nil {
		log.Fatalln("MySQL connection error: ", err)
	}

	pool, _ := db.DB()
	pool.SetMaxIdleConns(10)
	pool.SetMaxOpenConns(500)
	pool.SetConnMaxLifetime(time.Hour)

	log.Println("MySQL connection success")

	return db
}
