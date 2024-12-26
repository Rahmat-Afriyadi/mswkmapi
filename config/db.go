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

// func NewEnvConfig() []map[string]inteface{} {
// 	errEnv := godotenv.Load("../.env")
// 	if errEnv != nil {
// 		panic("Failed to load env file")
// 	}

// 	dbUser := os.Getenv("DB_USER")
// 	dbPass := os.Getenv("DB_PASS")
// 	dbHost := os.Getenv("DB_HOST")
// 	dbName := os.Getenv("DB_NAME")

// 	return []map[string]inteface{}{
// 		{
// 			"db_wkm":
// 		}
// 	}

// }

// func GetConnection() *sql.DB {
// 	errEnv := godotenv.Load()
// 	if errEnv != nil {
// 		fmt.Println("ini errornya ", errEnv)
// 		panic("Failed to load env file")
// 	}

// 	db, err := sql.Open("mysql", os.Getenv("DB_WKM"))
// 	// db, err := sql.Open("mysql", "root2:root2@tcp(192.168.70.30:3306)/db_wkm?parseTime=true")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	db.SetMaxIdleConns(5)
// 	db.SetMaxOpenConns(20)
// 	db.SetConnMaxIdleTime(5 * time.Minute)
// 	db.SetConnMaxLifetime(60 * time.Minute)

// 	return db
// }

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

// func NewAsuransiGorm() (*gorm.DB, *sql.DB) {
// 	errEnv := godotenv.Load()
// 	if errEnv != nil {
// 		fmt.Println("ini errornya ", errEnv)
// 		panic("Failed to load env file")
// 	}
// 	time.LoadLocation("Asia/Jakarta")
// 	dsn := os.Getenv("WANDA_ASURANSI")
// 	// dsn := "root2:root2@tcp(192.168.12.171:3306)/wanda_asuransi?parseTime=true&loc=Asia%2FJakarta"

// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
// 		Logger:                                   logger.Default.LogMode(logger.Info),
// 		SkipDefaultTransaction:                   true,
// 		DisableForeignKeyConstraintWhenMigrating: true,
// 		NamingStrategy: schema.NamingStrategy{
// 			NoLowerCase:         true,
// 			IdentifierMaxLength: 30,
// 		},
// 		PrepareStmt:     true,
// 		CreateBatchSize: 50,
// 	})
// 	if err != nil {
// 		fmt.Println("Masuk sini gk guys logger ", err)
// 		panic(err)
// 	}
// 	sqlDB, _ := db.DB()

// 	sqlDB.SetMaxOpenConns(50)
// 	sqlDB.SetMaxIdleConns(10)
// 	sqlDB.SetConnMaxLifetime(20 * time.Minute)
// 	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
// 	return db, sqlDB
// }
