package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	user, ok := os.LookupEnv("DB_USER")
	if !ok {
		user = "root"
	}

	pass, ok := os.LookupEnv("DB_PASS")
	if !ok {
		pass = "secret"
	}

	host, ok := os.LookupEnv("DB_HOST")
	if !ok {
		host = "localhost"
	}

	port, ok := os.LookupEnv("DB_PORT")
	if !ok {
		port = "3306"
	}

	name, ok := os.LookupEnv("DB_NAME")
	if !ok {
		name = "shapes"
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to open db connection: \n%+v\n", err)
	}

	sql, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql-db : \n%+v\n", err)
	}

	if err := sql.Ping(); err != nil {
		log.Fatalf("failed to ping db connection: \n%+v\n", err)
	}

	return db
}

func Close(db *gorm.DB) error {
	sql, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql-db : %w", err)
	}
	return sql.Close()
}
