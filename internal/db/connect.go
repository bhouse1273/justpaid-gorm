package db

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func MustOpen() *gorm.DB {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		log.Fatal("MYSQL_DSN is not set")
	}
	cfg := &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // our tables are singular
		},
		Logger: logger.Default.LogMode(logger.Warn),
	}
	db, err := gorm.Open(mysql.Open(dsn), cfg)
	if err != nil {
		log.Fatalf("open mysql: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("db.DB(): %v", err)
	}
	sqlDB.SetMaxOpenConns(40)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(2 * time.Hour)
	return db
}
