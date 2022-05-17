package model

type Order struct {
	Order_uid          string   `json:"order_uid"`
	Track_number       string   `json:"track_number"`
	Entry              string   `json:"entry"`
	Delivery           Delivery `json:"delivery"`
	Payment            Payment  `json:"payment"`
	Items              []Item   `json:"items"`
	Locale             string   `json:"locale"`
	Internal_signature string   `json:"internal_signature"`
	Customer_id        string   `json:"customer_id"`
	Delivery_service   string   `json:"delivery_service"`
	Shardkey           string   `json:"shardkey"`
	Sm_id              uint     `json:"sm_id"`
	Date_created       string   `json:"date_created"`
	Oof_shard          string   `json:"oof_shard"`
}

type Payment struct {
	Transaction   string `json:"transaction"`
	Request_id    string `json:"request_id"`
	Currency      string `json:"currency"`
	Provider      string `json:"provider"`
	Amount        uint   `json:"amount"`
	Payment_dt    uint   `json:"payment_dt"`
	Bank          string `json:"bank"`
	Delivery_cost uint   `json:"delivery_cost"`
	Goods_total   uint   `json:"goods_total"`
	Custom_fee    uint   `json:"custom_fee"`
}

type Item struct {
	Chrt_id      uint   `json:"chrt_id"`
	Track_number string `json:"track_number"`
	Price        uint   `json:"price"`
	Rid          string `json:"rid"`
	Name         string `json:"name"`
	Sale         uint   `json:"sale"`
	Size         string `json:"size"`
	Total_price  uint   `json:"total_price"`
	Nm_id        uint   `json:"nm_id"`
	Brand        string `json:"brand"`
	Status       uint   `json:"status"`
}

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
	Id      string `json:"id"`
}
