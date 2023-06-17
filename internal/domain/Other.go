package domain

import (

	"gorm.io/gorm"
	"time"
)

var Db *gorm.DB

type Channel struct {
	Id          uint        `gorm:"primarykey;autoIncrement;not null"`
	OwnerId     uint        `gorm:"not null; foreignKey:id onUpdate:CASCADE; onDelete:CASCADE"`
	Owner       UserModel   `json:"-"`
	Name        string      `gorm:"type:varchar(255);"`
	Description string      `gorm:"type:varchar(255);"`
	SocialLink  string      `gorm:"type:varchar(255);"`
	Banner      string      `gorm:"type:varchar(255);"`
	Icon        string      `gorm:"type:varchar(255);"`
	Subscribers []UserModel `gorm:"many2many:channel_subscription;"`
}

type Role struct {
	Id          uint `gorm:"primarykey;autoIncrement;not null"`
	ChannelId   int
	Channel     Channel     `gorm:"foreignKey:channel_id; onUpdate:CASCADE; onDelete:CASCADE"`
	Users       []UserModel `gorm:"many2many:user_roles; onUpdate:CASCADE; onDelete:CASCADE"`
	Weight      int         `gorm:"integer;"`
	Permission  int64       `gorm:"type:bigint"`
	Name        string      `gorm:"type:varchar(255);"`
	Description string      `gorm:"type:varchar(255);"`
}

type UserModel struct {
	Id            uint      			`gorm:"primarykey;autoIncrement;not null"`
	Icon          string    			`gorm:"type:varchar(255);"`
	Username      string    			`gorm:"type:varchar(255);not null"`
	Email         string    			`gorm:"type:varchar(255);"`
	Password      string    			`gorm:"type:varchar(255);"`
	Permission    int64     			`gorm:"type:bigint;default:1380863"`
	Incredentials string    			`gorm:"column:credentials type:text"`
	ValideAccount bool      			`gorm:"type:bool; default false"`
	Disable       bool      			`gorm:"type:bool; default false"`
	Online        bool      			`gorm:"type:bool; default false"`
	Subscribtion  []Channel 			`gorm:"many2many:channel_subscription;  onUpdate:CASCADE; onDelete:CASCADE"`
	Roles         []Role    			`gorm:"many2many:user_roles; onUpdate:CASCADE; onDelete:CASCADE"`
	CreatedAt     time.Time             `gorm:"default:CURRENT_TIMESTAMP"`
}

type Videos struct {
	Id            uint   `gorm:"primarykey;autoIncrement;not null"`
	Name          string `gorm:"type:varchar(255);"`
	Description   string `gorm:"type:varchar(1500);"`
	Icon          string `gorm:"type:varchar(255);"`
	VideoURL      string `gorm:"type:varchar(255);"`
	Views         int    `gorm:"type:integer"`
	Size          int64  `gorm:"type:integer"`
	ChannelId     uint   `gorm:"foreignKey:id"`
	Channel       Channel
	CreatedAt     string `gorm:"column:created_at"`
	CreationDate  string `gorm:"column:creation_date"`
	IsBlock       bool   `gorm:"type:boolean;default:false"`
	IsHide        bool   `gorm:"type:boolean;default:false"`
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

func (user *UserModel) Get() (*UserModel, error) {
	err := Db.Where("id = ?", user.Id).First(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (video *Videos) Get() (*Videos, error) {
	err := Db.Where("id = ?", video.Id).First(video).Error

	if err != nil {
		return nil, err
	}

	return video, nil
}