package gorm

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	pomodork_constant "pomodork-backend/constant"
	pomodork_error "pomodork-backend/message/error"
	pomodork_util "pomodork-backend/util"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GormHandler Gormのハンドラー
type GormHandler struct {
	db *gorm.DB
}

// NewGormHandler GormHandlerの生成
func NewGormHandler() (*GormHandler, error) {

	dsnFormat := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local"

	// DB接続情報取得
	dbUser, err := pomodork_util.GetRequiredEnv(pomodork_constant.DB_USER, "string")
	if err != nil {
		return nil, err
	}
	dbPassword, err := pomodork_util.GetRequiredEnv(pomodork_constant.DB_PASSWORD, "string")
	if err != nil {
		return nil, err
	}
	dbAddress, err := pomodork_util.GetRequiredEnv(pomodork_constant.DB_ADDRESS, "string")
	if err != nil {
		return nil, err
	}
	dbPort, err := pomodork_util.GetRequiredEnv(pomodork_constant.DB_PORT, "string")
	if err != nil {
		return nil, err
	}
	dbName, err := pomodork_util.GetRequiredEnv(pomodork_constant.DB_NAME, "string")
	if err != nil {
		return nil, err
	}

	// DBコネクション情報
	maxOpenConn, err := pomodork_util.GetRequiredEnv(pomodork_constant.MAX_OPEN_CONNECTION, "int")
	if err != nil {
		return nil, err
	}
	maxIdleConn, err := pomodork_util.GetRequiredEnv(pomodork_constant.MAX_IDLE_CONNECTION, "int")
	if err != nil {
		return nil, err
	}
	maxLifetime, err := pomodork_util.GetRequiredEnv(pomodork_constant.MAX_LIFE_TIME, "int")
	if err != nil {
		return nil, err
	}

	// DBコネクションプールの設定
	sqlDB := &sql.DB{}
	sqlDB.SetMaxOpenConns(maxOpenConn.(int))
	sqlDB.SetMaxIdleConns(maxIdleConn.(int))
	sqlDB.SetConnMaxLifetime(time.Duration(maxLifetime.(int)) * time.Millisecond)

	// Loggerの作成
	debugMode, _ := pomodork_util.GetNonRequiredEnv(pomodork_constant.DEBUG_MODE, "string", "0")
	logLevel := logger.Error
	if debugMode == "1" {
		logLevel = logger.Error
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	dsn := fmt.Sprintf(dsnFormat, dbUser, dbPassword, dbAddress, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:   newLogger,
		ConnPool: sqlDB,
	})
	if err != nil {
		logger := pomodork_util.NewLogger()
		dbErr := &pomodork_error.DBConnectError{
			Host: dbAddress.(string),
			Db:   dbName.(string),
			Port: dbPort.(string),
		}
		logger.Error(context.Background(), dbErr.Code(), dbErr.Error())
		return nil, dbErr
	}

	return &GormHandler{db: db}, nil
}

// Create レコード作成
func (gorm *GormHandler) Create(ctx context.Context, model interface{}) error {
	tx := gorm.db.Create(model)
	return tx.Error
}

// Find レコード取得
func (gorm *GormHandler) Find(ctx context.Context, model interface{}) error {
	tx := gorm.db.Find(model)
	return tx.Error
}

// Latest 最新のレコード取得
func (gorm *GormHandler) Latest(ctx context.Context, model interface{}) error {
	tx := gorm.db.Last(model)
	return tx.Error
}

// Count レコード数取得
func (gorm *GormHandler) Count(ctx context.Context, model interface{}) (int64, error) {
	var count int64
	tx := gorm.db.Find(model).Count(&count)
	return count, tx.Error
}
