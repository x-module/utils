/**
* Created by GoLand
* @file mysql.go
* @version: 1.0.0
* @author 李锦 <Lijin@cavemanstudio.net>
* @date 2022/1/27 11:42 上午
* @desc 初始化管理后台数据库
 */

package dirver

import (
	"fmt"
	"github.com/go-xmodule/utils/global"
	"github.com/go-xmodule/utils/utils/xerror"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

const (
	// DebugMode indicates gin mode is debug.
	DebugMode = "debug"
	DbType    = "mysql"
)

var DB *gorm.DB

type LinkParams struct {
	Host        string
	Port        int
	UserName    string
	DbName      string
	Password    string
	MaxOpenConn int
	MaxIdleConn int
	Mode        string
}

// InitializeDB 初始化管理后台数据库
func InitializeDB(params LinkParams) (*gorm.DB, error) {
	linkParams := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	dsn := fmt.Sprintf(linkParams, params.UserName, params.Password, params.Host, params.Port, params.DbName)
	fmt.Println(dsn)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	fmt.Println(err)
	fmt.Println("--------------------------")
	xerror.PanicErr(err, global.ConnectMysqlErr.String())
	// 链接池设置
	// db.DB().SetMaxOpenConns(params.MaxOpenConn)
	// db.DB().SetMaxIdleConns(params.MaxIdleConn)
	// db.LogMode(params.Mode == DebugMode)
	// db.LogMode(false)
	return db, nil
}
