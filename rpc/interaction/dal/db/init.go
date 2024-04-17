package db

import (
	"bibi/pkg/constants"
	"bibi/pkg/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	DBLike        *gorm.DB
	DBComment     *gorm.DB
	DBCommentLike *gorm.DB
)

func Init() {
	var err error
	DBLike, err = gorm.Open(mysql.Open(utils.InitMysqlDSN()),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, //单数表名
			},
		})
	if err != nil {
		klog.Fatal("db_like connect error")
	} else {
		klog.Info("db_like connect access")
	}
	sqlDB, _ := DBLike.DB()
	sqlDB.SetMaxIdleConns(constants.MaxIdleConns)
	sqlDB.SetMaxOpenConns(constants.MaxConnections)
	sqlDB.SetConnMaxLifetime(constants.ConnMaxLifetime)
	DBLike = DBLike.Table(constants.LikeTableName)

	DBComment, err = gorm.Open(mysql.Open(utils.InitMysqlDSN()),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, //单数表名
			},
		})
	if err != nil {
		klog.Fatal("db_comment connect error")
	} else {
		klog.Info("db_comment connect access")
	}
	//sqlDB, err = DBCommentLike.DB()
	//sqlDB.SetMaxIdleConns(constants.MaxIdleConns)
	//sqlDB.SetMaxOpenConns(constants.MaxConnections)
	//sqlDB.SetConnMaxLifetime(constants.ConnMaxLifetime)
	DBComment = DBComment.Table(constants.CommentTableName)

	DBCommentLike, err = gorm.Open(mysql.Open(utils.InitMysqlDSN()),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, //单数表名
			},
		})
	if err != nil {
		klog.Fatal("db_comment_like connect error")
	} else {
		klog.Info("db_comment_like connect access")
	}
	//sqlDB, err = DBCommentLike.DB()
	//sqlDB.SetMaxIdleConns(constants.MaxIdleConns)
	//sqlDB.SetMaxOpenConns(constants.MaxConnections)
	//sqlDB.SetConnMaxLifetime(constants.ConnMaxLifetime)
	DBCommentLike = DBCommentLike.Table(constants.CommentLikeTableName)
}
