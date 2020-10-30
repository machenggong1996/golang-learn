package models

import "time"

type Topic struct {
	TopicID    int    `json:"topicId"`
	TopicTitle string `json:"topicTitle" binding:"required"`
	TopicDate time.Time
	TopicUrl   string `json:"url" binding:"omitempty,topicUrl"`
}

type TopicClass struct {
	ClassId int `gorm:"PRIMARY_KEY"`
	ClassName string
	ClassRemark string
	ClassType string `gorm:"Column:classtype"`
}