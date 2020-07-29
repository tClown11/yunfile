package comm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func init() {
	db, err = gorm.Open("mysql", "root:Tj121108.96@/yunfile?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	//设置数据库连接池参数
	db.DB().SetMaxOpenConns(100) //设置数据库连接池最大连接数
	db.DB().SetMaxIdleConns(10)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
}

func GetDB() *gorm.DB {
	return db
}
