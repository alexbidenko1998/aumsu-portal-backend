package models

import (
	"aumsu.portal.backend/dif"
	"aumsu.portal.backend/entities"
	"gorm.io/gorm"
)

type MessageModel struct {
}

func (messageModel MessageModel) GetLast() (entities.Message, error) {
	var message entities.Message
	err := dif.DB.Model(&entities.Message{}).Last(&message).Error

	if err != nil {
		return message, err
	}

	return message, nil
}

func (messageModel MessageModel) Create(model *entities.Message) {
	dif.DB.Model(&entities.Message{}).Create(model)
}

func (messageModel MessageModel) All() []entities.Message {
	var messages []entities.Message
	dif.DB.Model(&entities.Message{}).Limit(40).Order("created_at desc").Order("id asc").Find(&messages)
	return messages
}

func (messageModel MessageModel) GetById(id string) (entities.Message, error) {
	var message entities.Message
	err := dif.DB.Model(&entities.Message{}).Preload("Comments", func(db *gorm.DB) *gorm.DB {
		return db.Order("comments.created_at ASC").Order("comments.id ASC")
	}).Preload("Comments.User").First(&message, id).Error

	if err != nil {
		return message, err
	}

	return message, nil
}
