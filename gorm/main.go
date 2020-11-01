package main

import (
	"fmt"
	. "github.com/machenggong1996/golang-learn/gorm/dbinit"
	. "github.com/machenggong1996/golang-learn/gorm/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
)

func main() {
	//Select()
	//Create()
	//CreateBatch()
	//First()
	//DryRunSelect()
	//TxUse()
	//TxRollBackPart()
	TopicClassApiUse()
}

//创建
func Create() {
	topicClass := TopicClass{ClassName: `class-1`, ClassRemark: `remark-1`, ClassType: `type-1`}
	result := DbHelper.Create(&topicClass)
	log.Println(topicClass.ClassId)
	log.Println(result.RowsAffected)
	log.Println(result.Error)
}

//批量插入
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
	DbHelper.Debug().Table("topics").Where("topic_id=?", 1).Find(&tc)
	fmt.Println(tc)
}

func First() {
	tc := Topic{}
	// 获取第一条记录（主键升序）
	//DbHelper.Table("topics").Take(&tc)
	// SELECT * FROM users ORDER BY id LIMIT 1;
	DbHelper.Table("topics").Last(&tc)
	fmt.Println(tc)
}

func DryRunSelect() {
	tc := Topic{}
	stmt := DbHelper.Session(&gorm.Session{DryRun: true}).First(&tc, 1).Statement
	fmt.Println(stmt.SQL.String()) //=> SELECT * FROM `users` WHERE `id` = $1 ORDER BY `id`
	fmt.Println(stmt.Vars)         //=> []interface{}{1}
}

//使用事务
func TxUse() {
	DbHelper.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Create(&TopicClass{ClassName: "Giraffe-err"}).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		if err := tx.Create(&TopicClass{ClassName: "Lion-err"}).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return errors.New("roll back")
	})
}

//嵌套事务 回滚部分事务 mysql SavePoint
func TxRollBackPart() {

	topicClass1 := TopicClass{ClassName: "topic1"}
	topicClass2 := TopicClass{ClassName: "topic2"}
	topicClass3 := TopicClass{ClassName: "topic3"}

	DbHelper.Transaction(func(tx *gorm.DB) error {
		tx.Create(&topicClass1)

		tx.Transaction(func(tx2 *gorm.DB) error {
			tx2.Create(&topicClass2)
			return errors.New("rollback topic") // 回滚
		})

		tx.Transaction(func(tx2 *gorm.DB) error {
			tx2.Create(&topicClass3)
			return nil
		})

		return nil
	})
}

//Api查询使用 查询部分数据
func TopicClassApiUse(){
	var api []ApiTopicClass
	DbHelper.Model(&TopicClass{}).Limit(10).Find(&api)
	fmt.Println(api)
}
