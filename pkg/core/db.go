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

type UserCache struct {
	users map[int64]*User
}

func NewUserCache() *UserCache {
	return &UserCache{
		users: make(map[int64]*User),
	}
}

func (mu *UserCache) Get(uid int64) *User {
	u := mu.users[uid]
	if u == nil {
		u = new(User)
		ret := DB.Take(&u, "id=?", uid)
		if ret.RowsAffected != 1 {
			u.Name = "Not Found"
			u.Avatar = CFG.Site.DefaultAvatar
		}
		if len(u.Avatar) == 0 {
			u.Avatar = CFG.Site.DefaultAvatar
		}
		mu.users[uid] = u
	}
	return u
}

type User struct {
	ID          int64  `gorm:"primaryKey"`
	Name        string `gorm:"column:u_name"`
	Avatar      string
	Likes       int64
	Email       string
	Passwd      string
	Phone       string
	CountryCode int
	CreateTime  time.Time
	Sha1        string
	ShortSha1   string `gorm:"column:sha1_prefix"`
	Valid       int
	OnceToken   string
}

func (User) TableName() string {
	return "t_user"
}

type Region struct {
	ID         int64  `gorm:"primaryKey"`
	Name       string `gorm:"column:r_name"`
	About      string
	CreateTime time.Time
	Sha1       string
	ShortSha1  string `gorm:"column:sha1_prefix"`
}

func (Region) TableName() string {
	return "t_region"
}

type Discuss struct {
	ID           int64  `gorm:"primaryKey"`
	Name         string `gorm:"column:d_name"`
	Content      string
	InitiatorUid int64
	Initiator    string
	Likes        int64
	DivisionRid  int64
	Division     string
	CreateTime   time.Time
	Sha1         string
	ShortSha1    string `gorm:"column:sha1_prefix"`
}

func (Discuss) TableName() string {
	return "t_discuss"
}

type Comment struct {
	ID           int64 `gorm:"primaryKey"`
	Content      string
	InitiatorUid int64
	Initiator    string
	Likes        int64
	DiscussDid   int64
	CreateTime   time.Time
	Sha1         string
	ShortSha1    string `gorm:"column:sha1_prefix"`
}

func (Comment) TableName() string {
	return "t_comment"
}
