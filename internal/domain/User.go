package domain

import (

	"gorm.io/gorm"
)

var Db *gorm.DB

type Channel struct {
	Id          uint `gorm:"primarykey;autoIncrement;not null"`
	OwnerId     uint `gorm:"not null; foreignKey:Owner"`
	Owner       UserModel
	Description string      `gorm:"type:varchar(255);"`
	SocialLink  string      `gorm:"type:varchar(255);"`
	Banner      string      `gorm:"type:varchar(255);"`
	Icon        string      `gorm:"type:varchar(255);"`
	Subscribers []UserModel `gorm:"many2many:channel_subscription;"`
}

type Role struct {
	Id          uint `gorm:"primarykey;autoIncrement;not null"`
	ChannelId   int
	Channel     Channel     `gorm:"foreignKey:ChannelId"`
	User        []UserModel `gorm:"many2many:user_roles;"`
	Permission  int64       `gorm:"type:bigint"`
	Name        string      `gorm:"type:varchar(255);"`
	Description string      `gorm:"type:varchar(255);"`
}

type UserModel struct {
	Id            uint      `gorm:"primarykey;autoIncrement;not null"`
	Icon          string    `gorm:"type:varchar(255);"`
	Username      string    `gorm:"type:varchar(255);not null"`
	Email         string    `gorm:"type:varchar(255);"`
	Password      string    `gorm:"type:varchar(255);"`
	Permission    int64     `gorm:"type:bigint"`
	Incredentials string    `gorm:"column:credentials type:text"`
	ValideAccount bool      `gorm:"type:bool; default false"`
	Disable       bool      `gorm:"type:bool; default false"`
	Subscribtion  []Channel `gorm:"many2many:channel_subscription;"`
	Role          []Role    `gorm:"many2many:user_roles;"`
}

type Videos struct {
	Id          uint   `gorm:"primarykey;autoIncrement;not null"`
	Name        string `gorm:"type:varchar(255);"`
	Description string `gorm:"type:varchar(255);"`
	Icon        string `gorm:"type:varchar(255);"`
	VideoURL    string `gorm:"type:varchar(255);"`
	ChannelId   uint   `gorm:"foreignKey:id"`
	Channel     Channel
	CreatedAt   string `gorm:"type:time without time zone"`
}

func (user *UserModel) TableName() string {
	return "users"
}
func (r *Role) TableName() string {
	return "roles"
}
func (channel *Channel) TableName() string {
	return "channels"
}

func (video *Videos) TableName() string {
	return "video_info"
}
