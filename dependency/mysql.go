package driver

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SQLConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DB       string
}

func NewRegistrySQLWithORM(cfg SQLConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DB, "parseTime=true")

	sqlDB, errConnect := sql.Open("mysql", dsn)
	if errConnect != nil {
		return nil, errConnect
	}
	sqlDB.SetConnMaxIdleTime(10)
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}))
	if err != nil {
		return db, err
	}

	return db, nil
}
