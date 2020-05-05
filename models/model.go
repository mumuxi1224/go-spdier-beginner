package models

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type DatabaseInfo struct {
	User,Password,Host,Name,TablePrefix,Type string
}

var (
	Db *gorm.DB
	databaseInfo = &DatabaseInfo{}
	err error
)

func init() {
	Cfg,err:=ini.Load("conf/config.ini")
	if err!=nil {
		log.Fatalf("load config error:%s",err)
	}
	err= Cfg.Section("database").MapTo(databaseInfo)
	if err!=nil {
		log.Fatalf("get database configure error：%s",err)
	}
	Db, err = gorm.Open(databaseInfo.Type, fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		databaseInfo.User,
		databaseInfo.Password,
		databaseInfo.Host,
		databaseInfo.Name,

		))

	if err!=nil {
		log.Fatalf("database connection error:%s",err)
	}
	err=Db.DB().Ping()
	if err!=nil {
		log.Fatalf("connection to the database is not  alive :%s",err)
	}


}

func ClosrDb() {
	defer Db.Close()
}
