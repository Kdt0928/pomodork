package gorm

import (
	"context"
	"database/sql"
	"fmt"
	pomodork_constant "pomodork-backend/constant"
	"pomodork-backend/interface/database"
	pomodork_util "pomodork-backend/util"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormHandler struct {
	db *gorm.DB
}

// TODO Loggerを作る

// NewGormHandler gormHandlerの生成
func NewGormHandler() (*GormHandler, error) {

	// logger := pomodork_util.NewLogger()
	// dsnFormat := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=%s"
	dsnFormat := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local"

	// DB接続情報取得
	dbUser, err := pomodork_util.GetEnv(pomodork_constant.DB_USER, "string")
	if err != nil {
		return nil, err
	}
	dbPassword, err := pomodork_util.GetEnv(pomodork_constant.DB_PASSWORD, "string")
	if err != nil {
		return nil, err
	}
	dbAddress, err := pomodork_util.GetEnv(pomodork_constant.DB_ADDRESS, "string")
	if err != nil {
		return nil, err
	}
	dbPort, err := pomodork_util.GetEnv(pomodork_constant.DB_PORT, "string")
	if err != nil {
		return nil, err
	}
	dbName, err := pomodork_util.GetEnv(pomodork_constant.DB_NAME, "string")
	if err != nil {
		return nil, err
	}
	// dbTimeZone, err := pomodork_util.GetEnv(pomodork_constant.DB_TIMEZONE, "string")
	// if err != nil {
	// 	return nil, err
	// }

	// DBコネクション情報
	maxOpenConn, err := pomodork_util.GetEnv(pomodork_constant.MAX_OPEN_CONNECTION, "int")
	if err != nil {
		return nil, err
	}
	maxIdleConn, err := pomodork_util.GetEnv(pomodork_constant.MAX_IDLE_CONNECTION, "int")
	if err != nil {
		return nil, err
	}
	maxLifetime, err := pomodork_util.GetEnv(pomodork_constant.MAX_LIFE_TIME, "int")
	if err != nil {
		return nil, err
	}

	// DBコネクションプールの設定
	sqlDB := &sql.DB{}
	sqlDB.SetMaxOpenConns(maxOpenConn.(int))
	sqlDB.SetMaxIdleConns(maxIdleConn.(int))
	sqlDB.SetConnMaxLifetime(time.Duration(maxLifetime.(int)) * time.Millisecond)

	// // TODO Loggerの作成
	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags),
	// 	logger.Config{
	// 		SlowThreshold:             time.Second,
	// 		LogLevel:                  logger.Info,
	// 		IgnoreRecordNotFoundError: true,
	// 		Colorful:                  true,
	// 	},
	// )

	// dsn := fmt.Sprintf(dsnFormat, dbUser, dbPassword, dbAddress, dbPort, dbName, dbTimeZone)
	dsn := fmt.Sprintf(dsnFormat, dbUser, dbPassword, dbAddress, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger:   newLogger,
		ConnPool: sqlDB,
	})
	if err != nil {
		// logger.Error()
		return nil, err // TODO ログ出力
	}

	return &GormHandler{db: db}, nil
}

// GetMaxUserId 最新のユーザID取得
func (gorm *GormHandler) GetMaxUserId(ctx context.Context, userAccount *database.UserAccount) error {
	tx := gorm.db.Last(userAccount)
	return tx.Error
}

// UserCount ユーザ数カウント
func (gorm *GormHandler) UserCount(ctx context.Context, userAccount *database.UserAccount) (int64, error) {
	var count int64
	tx := gorm.db.Find(userAccount).Count(&count)
	return count, tx.Error
}

// CreateUser ユーザ作成
func (gorm *GormHandler) CreateUser(ctx context.Context, userAccount *database.UserAccount) error {
	tx := gorm.db.Create(userAccount)
	return tx.Error
}
