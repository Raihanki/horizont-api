package configs

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectDatabase() (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		ENV.DB_USERNAME,
		ENV.DB_PASSWORD,
		ENV.DB_HOST,
		ENV.DB_PORT,
		ENV.DB_DATABASE,
	)

	db, errDb := sqlx.Open("mysql", dsn)
	if errDb != nil {
		return nil, errDb
	}

	errPing := db.Ping()
	if errPing != nil {
		return nil, errPing
	}

	db.SetConnMaxLifetime(time.Duration(ENV.DB_MAX_LIFETIME_CONNECTION) * time.Second)
	db.SetConnMaxIdleTime(time.Duration(ENV.DB_MAX_IDLE_TIME_CONNECTION) * time.Second)
	db.SetMaxIdleConns(int(ENV.DB_MAX_IDLE_CONNECTION))
	db.SetMaxOpenConns(int(ENV.DB_MAX_OPEN_CONNECTION))

	log.Info("Database connected")

	return db, nil
}
