package gorm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var Client *GormClient

type GormClient struct {
	dataSourceName string
	logMode        bool
}

func InitialDBClient(dataSourceName string, logMode bool) error {
	log.Infof("initial db client. dataSourceName : %v", dataSourceName)
	//db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}
	db.SingularTable(true) //禁用复数表名
	db.LogMode(logMode)
	Client = &GormClient{dataSourceName: dataSourceName, logMode: logMode}
	defer db.Close()
	return nil
}

func (gc GormClient) GetDBClient() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", gc.dataSourceName)
	if err != nil {
		return nil, err
	}
	db.LogMode(gc.logMode)
	db.SingularTable(true)
	return db, nil
}
