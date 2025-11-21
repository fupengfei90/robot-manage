package database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/example/robot-manage/internal/config"
)

// New 初始化数据库连接（默认MySQL），并配置连接池。
func New(cfg config.DatabaseSection) (*gorm.DB, error) {
	if cfg.Driver == "" || cfg.DSN == "" {
		return nil, fmt.Errorf("数据库配置缺失")
	}

	var dialector gorm.Dialector
	switch cfg.Driver {
	case "mysql", "mariadb":
		dialector = mysql.New(mysql.Config{
			DSN:                       cfg.DSN,
			DefaultStringSize:         256,
			SkipInitializeWithVersion: true,
		})
	default:
		return nil, fmt.Errorf("暂不支持的数据库驱动: %s", cfg.Driver)
	}

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	return db, nil
}
