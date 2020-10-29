package main

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/hashicorp/go-uuid"
	"strconv"
	"time"

	//"github.com/tidwall/gjson"
)

var Address = []string{"10.10.0.142:9092"}

func main() {
	syncProducer(Address)
	//asyncProducer1(Address)
}

//同步消息模式
func syncProducer(address []string) {
	fmt.Printf("producer_test\n")
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V0_11_0_2

	producer, err := sarama.NewAsyncProducer(address, config)
	if err != nil {
		fmt.Printf("producer_test create producer error :%s\n", err.Error())
		return
	}

	defer producer.AsyncClose()

	// send message
	msg := &sarama.ProducerMessage{
		Topic: "new-shop-shuhaisc-hdl-pos",
		//Key:   sarama.StringEncoder("go_test"),
	}
	//var message = `{"data":[{"id":"%s","shop_id":"BJ16","table_bill_id":"XXXXXX1","dish_code":"01060228","dish_id":"01060228","dish_name":"辣炫风火锅","dish_price":"100.0","dish_abnormal_status":"","order_time":"1583934415000","served_quantity":"400","dish_type":"3","standard_id":"半份","standard_code":"01","ts":"1583934415000"}],"database":"saptogaia","es":1583402646000,"id":5,"isDdl":false,"mysqlType":{"id":"char(36)","shop_id":"char(36)","table_bill_id":"char(36)","dish_code":"varchar(50)","dish_id":"char(36)","dish_name":"varchar(50)","dish_price":"varchar(20)","dish_abnormal_status":"varchar(30)","order_time":"bigint(20)","served_quantity":"varchar(11)","dish_type":"smallint(6)","standard_id":"char(36)","standard_code":"varchar(30)","ts":"bigint(20)","create_time":"timestamp(0)","last_modify_time":"timestamp(0)"},"old":null,"pkNames":["id"],"sql":"","sqlType":{"id":1,"shop_id":1,"table_bill_id":1,"dish_code":12,"dish_id":1,"dish_name":12,"dish_price":12,"dish_abnormal_status":12,"order_time":-5,"served_quantity":12,"dish_type":5,"standard_id":1,"standard_code":12,"ts":-5,"create_time":93,"last_modify_time":93},"table":"t_pro_dish_list_detail","ts":1583934415000,"type":"INSERT"}`

	var message = `{"data":%s,"database":"saptogaia","es":1583402646000,"id":5,"isDdl":false,"mysqlType":{"id":"char(36)","shop_id":"char(36)","table_bill_id":"char(36)","dish_code":"varchar(50)","dish_id":"char(36)","dish_name":"varchar(50)","dish_price":"varchar(20)","dish_abnormal_status":"varchar(30)","order_time":"bigint(20)","served_quantity":"varchar(11)","dish_type":"smallint(6)","standard_id":"char(36)","standard_code":"varchar(30)","ts":"bigint(20)","create_time":"timestamp(0)","last_modify_time":"timestamp(0)"},"old":null,"pkNames":["id"],"sql":"","sqlType":{"id":1,"shop_id":1,"table_bill_id":1,"dish_code":12,"dish_id":1,"dish_name":12,"dish_price":12,"dish_abnormal_status":12,"order_time":-5,"served_quantity":12,"dish_type":5,"standard_id":1,"standard_code":12,"ts":-5,"create_time":93,"last_modify_time":93},"table":"t_pro_dish_list_detail","ts":1583934415000,"type":"INSERT"}`

	message_list := []string{}
	t := time.Now().UnixNano() / 1e6

	for i := 1; i <= 100; i++ {
		dishs := []Dish{}
		for j := 1; j <= 100; j++ {
			uid, _ := GetUUID()
			dish := Dish{
				Id:                 uid,
				ShopId:             "BJ16",
				TableBillId:        "XXXXXX1",
				DishCode:           "12010001",
				DishId:             "12010001",
				DishName:           "自选小料",
				DishE:              "100.0",
				DishAbnormalStatus: "",
				OrderTime:          "1583934415000",
				ServedQuantity:     "400",
				DishType:           strconv.Itoa(i),
				StandardId:         "半份",
				StandarCode:        "01",
				Ts:                 strconv.FormatInt(t, 10),
			}
			dishs = append(dishs, dish)
		}
		jb, _ := json.Marshal(dishs)
		js := string(jb)
		m := fmt.Sprintf(message, js)
		message_list = append(message_list, m)
	}

	for i := 0; i < len(message_list); i++ {
		//id := strconv.Itoa(i)
		//value := fmt.Sprintf(message, id)
		value := message_list[i]
		msg.Value = sarama.ByteEncoder(value)
		//fmt.Println(value)

		//send to chain
		producer.Input() <- msg

		select {
		case suc := <-producer.Successes():
			fmt.Printf("offset: %d,  timestamp: %s\n", suc.Offset, suc.Timestamp.String())
		case fail := <-producer.Errors():
			fmt.Printf("err: %s\n", fail.Err.Error())
		}
	}
}

func GetUUID() (u string, err error) {
	u, err = uuid.GenerateUUID()
	return
}
