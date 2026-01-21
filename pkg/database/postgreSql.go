package database

import (
	"ares/pkg/config"
	"ares/pkg/session"
	"context"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(session *session.Session, config config.Database) *gorm.DB {

	dbConn, err := dbOpen(config)
	if err != nil {
		session.LogError("DB Connection - ", err)
		err = nil
		if dbConn, err = dbOpen(config); err != nil {
			session.LogError("DB Connection Second Try - ", err)
			log.Fatal("Couldn't connect to postgreSql")
		}
	}
	session.LogInfo("PostgreSQL Connected")
	return dbConn
}

func dbOpen(config config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Host, config.Username, config.Password, config.Schema, config.Port, config.SslMode, config.Location)

	DbConn, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	sqlDB, err := DbConn.WithContext(ctx).DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get *sql.DB: %w", err)
	}

	sqlDB.SetConnMaxIdleTime(time.Duration(config.MinIdleConnections)) // batas waktu maksimum koneksi idle
	sqlDB.SetMaxOpenConns(config.MaxOpenConnections)                   // jumlah maksimum koneksi aktif

	return DbConn, err
}
