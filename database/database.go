package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gorm.io/gorm/logger"

)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
    dsn := fmt.Sprintf(
        "host=db  password=%s dbname=%s port=5439 sslmode=disable TimeZone=Asia/Shanghai",
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })

    if err != nil {
        log.Fatal("Failed to connect to database. \n", err)
        os.Exit(1)
    }

    log.Println("connected")
    db.Logger = logger.Default.LogMode(logger.Info)

    log.Println("running migrations")
    err = db.AutoMigrate(&models.Fact{})
    if err != nil {
        log.Fatal("Failed to run migrations. \n", err)
        os.Exit(1)
    }

    DB = Dbinstance{Db: db}
}
