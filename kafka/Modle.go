package main

type Dish struct {
	Id string `json:"id"`
	ShopId string `json:"shop_id"`
	TableBillId string `json:"table_bill_id"`
	DishCode string `json:"dish_code"`
	DishId string `json:"dish_id"`
    DishName string `json:"dish_name"`
    DishE string `json:"dish_e"`
	DishAbnormalStatus string `json:"dish_abnormal_status"`
	OrderTime string `json:"order_time"`
	ServedQuantity string `json:"served_quantity"`
	DishType string `json:"dish_type"`
	StandardId string `json:"standard_id"`
	StandarCode string `json:"standarcode"`
	Ts string `json:"ts"`
}
