package config

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)


func GetConnectionMain() (*gorm.DB, *sql.DB) {

	errEnv := godotenv.Load()
	if errEnv != nil {
		fmt.Println("ini errornya ", errEnv)
		panic("Failed to load env file")
	}
	// dsn := os.Getenv("MS_WKM")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Asia%sJakarta", db_user, db_pass, db_host, db_port, db_name, "%2F")
	time.LoadLocation("Asia/Jakarta")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase:         true,
			IdentifierMaxLength: 30,
		},
		PrepareStmt:     true,
		CreateBatchSize: 50,
	})
	if err != nil {
		fmt.Println("Error db users ", err)
		panic(err)
	}

	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	return db, sqlDB
}

func GetUserConnectionMain() (*gorm.DB, *sql.DB) {

	errEnv := godotenv.Load()
	if errEnv != nil {
		fmt.Println("ini errornya ", errEnv)
		panic("Failed to load env file")
	}
	// dsn := os.Getenv("MS_WKM")
	db_user := os.Getenv("DB_USER_USER")
	db_pass := os.Getenv("DB_USER_PASS")
	db_host := os.Getenv("DB_USER_HOST")
	db_port := os.Getenv("DB_USER_PORT")
	db_name := os.Getenv("DB_USER_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Asia%sJakarta", db_user, db_pass, db_host, db_port, db_name, "%2F")
	time.LoadLocation("Asia/Jakarta")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase:         true,
			IdentifierMaxLength: 30,
		},
		PrepareStmt:     true,
		CreateBatchSize: 50,
	})
	if err != nil {
		fmt.Println("Error db users ", err)
		panic(err)
	}

	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(10 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
	return db, sqlDB
}

