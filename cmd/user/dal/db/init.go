package db

import (
	"bibi/pkg/constants"
	"bibi/pkg/utils"
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/milvus-io/milvus-sdk-go/v2/client"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	DB        *gorm.DB
	MilvusCli client.Client
)

func InitMySQL() {
	var err error
	DB, err = gorm.Open(mysql.Open(utils.InitMysqlDSN()),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, //单数表名
			},
		})
	if err != nil {
		klog.Fatal("mysql connect error")
	} else {
		klog.Info("mysql connect access")
	}

	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(constants.MaxIdleConns)
	sqlDB.SetMaxOpenConns(constants.MaxConnections)
	sqlDB.SetConnMaxLifetime(constants.ConnMaxLifetime)
	DB = DB.Table(constants.UserTableName)
}
func InitMilvus() {
	cli, err := client.NewClient(context.Background(), client.Config{
		Address: "0.0.0.0:19530",
	})
	if err != nil {
		klog.Fatal("milvus connect error")
	}
	MilvusCli = cli
}
