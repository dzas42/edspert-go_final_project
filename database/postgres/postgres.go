package postgres

import (
	"final-project/internal/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var db *gorm.DB

func GetConnection() *gorm.DB {
	if db != nil {
		return db
	}
	log.Fatalf("connection has not been set, please initial connection")
	return nil
}

func InitialConnection(conf config.Config) error {
	var err error
	var loggerConf logger.Interface
	// setup logger
	if conf.IS_PRODUCTION {
		loggerConf = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second,   // Slow SQL threshold
				LogLevel:                  logger.Silent, // Log level
				IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
				Colorful:                  false,         // Disable color
			},
		)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", conf.DBHost, conf.DBUserName, conf.DBUserPassword, conf.DBName, conf.DBPort)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: loggerConf,
	})
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return err
		// control error
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return err
}
