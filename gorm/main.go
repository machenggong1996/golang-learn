package main

import (
	"fmt"
	. "github.com/machenggong1996/golang-learn/gorm/dbinit"
	. "github.com/machenggong1996/golang-learn/gorm/models"
	"log"
)

func main() {
	//Select()
	//Create()
	CreateBatch()
}

//创建
func Create() {
	topicClass := TopicClass{ClassName: `class-1`, ClassRemark: `remark-1`, ClassType: `type-1`}
	result := DbHelper.Create(&topicClass)
	log.Println(topicClass.ClassId)
	log.Println(result.RowsAffected)
	log.Println(result.Error)
}

//TODO 批量插入报错
func CreateBatch() {
	//var topicClasses = []TopicClass{{ClassName: "jinzhu1"}, {ClassName: "jinzhu2"}, {ClassName: "jinzhu3"}}
	//DbHelper.Create(&topicClasses)
	//for _, topicClass := range topicClasses {
	//	 // 1,2,3
	//	log.Println(topicClass.ClassId)
	//}
	//DbHelper.Model(&TopicClass{}).Create(map[string]interface{}{
	//	"ClassName": "jinzhu_1", "ClassRemark": "re11",
	//})

	DbHelper.Model(&TopicClass{}).Create([]map[string]interface{}{
		{"ClassName": "jinzhu_1", "ClassRemark": "re11"},
		{"ClassName": "jinzhu_2", "ClassRemark": "re12"},
	})
}

func Select() {
	tc := Topic{}
	DbHelper.Table("topics").Where("topic_id=?", 1).Find(&tc)
	fmt.Println(tc)
}
