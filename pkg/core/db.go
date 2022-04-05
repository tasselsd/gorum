package core

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadDatabase() {
	dsn := CFG.String("db.dsn")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
}

type User struct {
	ID         int64  `gorm:"primaryKey"`
	Name       string `gorm:"column:u_name"`
	Avatar     string
	Likes      string
	Email      string
	Phone      string
	ContryCode int
	CreateTime time.Time
	Sha1       string
	ShortSha1  string `gorm:"column:sha1_prefix"`
}
