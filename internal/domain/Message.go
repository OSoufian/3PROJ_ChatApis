package domain

import "time"

type LiveMessage struct {
	Id      uint   `gorm:"primarykey;autoIncrement;not null"`
	Content string `gorm:"type:varchar(255);"`
	VideoId uint   `gorm:"foreignKey:id"`
	Video   Videos
	UserId  uint `gorm:"foreignKey:id"`
	User    UserModel
	Created time.Time `json:"created"`
}

type Message struct {
	Id      uint   `gorm:"primarykey;autoIncrement;not null"`
	Content string `json:"content"`
	VideoId uint   `gorm:"foreignKey:id"`
	Video   Videos
	UserId  uint `gorm:"foreignKey:id"`
	User    UserModel
	Created time.Time `json:"created"`
}

func (msg *Message) TableName() string {
	return "messages"
}

func (msg *LiveMessage) TableName() string {
	return "live_messages"
}

func (msg *Message) GetById() *Message {
	tx := Db.Where("id = ?", msg.Id).Find(msg)
	if tx.RowsAffected == 0 {
		return nil
	}
	return msg
}

func (msg *Message) GetAll() ([]Message, error) {
	var results []Message
	err := Db.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (msg *Message) createMessage() error {
	tx := Db.Create(msg)

	return tx.Error
}

func (msg *Message) UpdateMessage() {
	Db.Save(&msg)
}

func (msg *Message) DeletMessage() {
	Db.Delete(msg)
}

func (msg *LiveMessage) GetById() *LiveMessage {
	tx := Db.Where("id = ?", msg.Id).Find(msg)
	if tx.RowsAffected == 0 {
		return nil
	}
	return msg
}

func (msg *LiveMessage) GetAll() ([]LiveMessage, error) {
	var results []LiveMessage
	err := Db.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (msg *LiveMessage) createMessage() error {
	tx := Db.Create(msg)

	return tx.Error
}

func (msg *LiveMessage) UpdateMessage() {
	Db.Save(&msg)
}

func (msg *LiveMessage) DeletMessage() {
	Db.Delete(msg)
}
