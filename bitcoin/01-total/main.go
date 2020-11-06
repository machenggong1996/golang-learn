package main

import "fmt"

func main() {
	//初始奖励
	reward := 50.0
	//每隔21万块 奖励衰减
	interval := 21.0//万

	total := 0.0

	for reward>0{
		amount:= reward * interval
		total = total + amount
		reward = reward * 0.5
		//fmt.Println(reward)
	}

	fmt.Println(total)





}
