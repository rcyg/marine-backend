package bootstrap

import (
	"fmt"
	stdlog "log"
	"marine-backend/cmd/flags"
	"marine-backend/internal/conf"
	"marine-backend/internal/db"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() {
	logLevel := logger.Silent
	if flags.Debug || flags.Dev {
		logLevel = logger.Info
	}
	newLogger := logger.New(
		stdlog.New(log.StandardLogger().Out, "\r\n", stdlog.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	gormConfig := &gorm.Config{
		Logger: newLogger,
	}
	var dB *gorm.DB
	var err error
	database := conf.Conf.Database
	dsn := database.DSN
	if dsn == "" {
		//[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=%s",
			database.User, database.Password, database.Host, database.Port, database.Name, database.SSLMode)
	}
	dB, err = gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		log.Fatalf("failed to connect database:%s", err.Error())
	}
	db.Init(dB)
}
