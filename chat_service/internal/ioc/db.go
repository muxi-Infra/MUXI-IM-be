package ioc

import (
	"github.com/muxi-Infra/MUXI-IM-be/chat_service/internal/repository/dao"
	"github.com/muxi-Infra/MUXI-IM-be/pkg/logger"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

func InitDB(l logger.Logger) *gorm.DB {
	type Config struct {
		DSN string `yaml:"dsn"`
	}
	var cfg Config
	if err := viper.UnmarshalKey("database", &cfg); err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{
		Logger: glogger.New(gormLoggerFunc(l.Debug), glogger.Config{
			SlowThreshold: 0,
			LogLevel:      glogger.Info, // Debug模式下打印Info级别日志
		}),
	})
	if err != nil {
		panic(err)
	}
	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}

type gormLoggerFunc func(msg string, fields ...logger.Field)

func (g gormLoggerFunc) Printf(s string, i ...interface{}) {
	g(s, logger.Field{Key: "args", Val: i})
}

func InitDBwithSqlite() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err = dao.InitTables(db); err != nil {
		panic(err)
	}
	return db
}
