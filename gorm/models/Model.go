package models

import (
	"gorm.io/gorm"
	"log"
	"time"
)

type Topic struct {
	TopicID    int    `json:"topicId"`
	TopicTitle string `json:"topicTitle" binding:"required"`
	TopicDate  time.Time
	TopicUrl   string `json:"url" binding:"omitempty,topicUrl"`
}

type TopicClass struct {
	ClassId     int `gorm:"PRIMARY_KEY"`
	ClassName   string
	ClassRemark string
	ClassType   string `gorm:"Column:classtype"`
}

type ApiTopicClass struct {
	ClassId   int
	ClassName string
}

func (u *TopicClass) BeforeCreate(tx *gorm.DB) (err error) {
	log.Println("BeforeCreate")
	return
}

func (u *TopicClass) AfterCreate(tx *gorm.DB) (err error) {
	log.Println("AfterCreate")
	return
}
