package infrastructure

import (
	"fmt"
	"github.com/ari1021/hack-ios-server/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewConnection は，migrationを行い，*gorm.DBを返します．
func NewConnection() (*gorm.DB, error) {
	DSN := config.DSN()
	conn, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open MySQL: %w", err)
	}
	//if err := migrate(conn); err != nil {
	//	log.Fatal(err)
	//}
	return conn, nil
}

// migrate は，migrationを行います．
//func migrate(conn *gorm.DB) error {
//	if err := conn.AutoMigrate(
//		&User{},
//		&Task{},
//	); err != nil {
//		return fmt.Errorf("failed to migrate: %w", err)
//	}
//	return nil
//}
