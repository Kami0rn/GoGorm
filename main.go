package main

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql,_ := fc()
	fmt.Printf("%v\n====================================\n",sql)
}

func main() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/tew?parseTime=true"
	dial := mysql.Open(dsn)
	db,err := gorm.Open(dial , &gorm.Config{
		Logger: &SqlLogger{},
		DryRun: true,
	})
	if err != nil {
		panic(err)
	}

	db.Migrator().CreateTable(Gender{})

}

type Gender struct {
	ID uint
	Name string
	CreatAt string `gorm:"column:myname"`
}