package mysql

import (
	"demo/internal/models/mysql/entity"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// Init ...
func Init() {
	openDB()
}

func openDB() (*gorm.DB, error) {
	var err error
	db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gin-demo?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		return db, err
	}
	setupDB(db)

	return db, nil
}

func setupDB(db *gorm.DB) {
	// 配置未debug 开启gorm debug
	// 用于设置最大打开的连接数，默认值为0表示不限制设置最大的连接数，
	// 可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxOpenConns(10)
	// 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	db.DB().SetMaxIdleConns(100)
	db.SingularTable(true)
	Auto(db)

}

func GetDB() *gorm.DB {
	return db
}

func Auto(db *gorm.DB) {
	db.AutoMigrate(&entity.User{})
}
