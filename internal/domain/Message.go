package domain

// import "time"

type LiveMessage struct {
	Message string `json:"message"`
}

// type LiveMessage struct {
// 	Message string `json:"message"`
// 	VideoId uint   `gorm:"foreignKey:id"`
// }

type Message struct {
	Id      uint       `gorm:"primarykey;autoIncrement;not null"`
	Content string     `json:"Content"`
	VideoId uint       `gorm:"foreignKey:id"`
	Video   Videos
	UserId  uint       `gorm:"foreignKey:id"`
	User    UserModel
	Created string     `gorm:"type:time without time zone"`
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

func (message *Message) GetAll(video int) ([]Message, error) {
	var results []Message
	err := Db.Where("video_id = ?", video).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (msg *Message) Create() error {
	tx := Db.Create(msg)

	return tx.Error
}

func (msg *Message) DeletMessage() {
	// Db.Where("id = ?", msg.Id).Find(msg)
	
	Db.Delete(msg)
}

// func (msg *LiveMessage) GetById() *LiveMessage {
// 	tx := Db.Where("id = ?", msg.Id).Find(msg)
// 	if tx.RowsAffected == 0 {
// 		return nil
// 	}
// 	return msg
// }

func (msg *LiveMessage) CreateMessage() error {
	tx := Db.Create(msg)

	return tx.Error
}