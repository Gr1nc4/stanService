package main

import (
	"github.com/nats-io/stan.go"
)

const (
	clusterName = "test-cluster"
	clientName  = "publisher"
	channel     = "orders"
)

func main() {
	sc, _ := stan.Connect(clusterName, clientName)
	defer sc.Close()
	so := `{
		"order_uid": "b563feb7b2b84b6test",
		"track_number": "WBILMTESTTRACK",
		"entry": "WBIL",
		"delivery": {
		  "name": "Test Testov",
		  "phone": "+9720000000",
		  "zip": "2639809",
		  "city": "Kiryat Mozkin",
		  "address": "Ploshad Mira 15",
		  "region": "Kraiot",
		  "email": "test@gmail.com"
		},
		"payment": {
		  "transaction": "b563feb7b2b84b6test",
		  "request_id": "",
		  "currency": "USD",
		  "provider": "wbpay",
		  "amount": 1817,
		  "payment_dt": 1637907727,
		  "bank": "alpha",
		  "delivery_cost": 1500,
		  "goods_total": 317,
		  "custom_fee": 0
		},
		"items": [
		  {
			"chrt_id": 9934930,
			"track_number": "WBILMTESTTRACK",
			"price": 453,
			"rid": "ab4219087a764ae0btest",
			"name": "Mascaras",
			"sale": 30,
			"size": "0",
			"total_price": 317,
			"nm_id": 2389212,
			"brand": "Vivienne Sabo",
			"status": 202
		  }
		],
		"locale": "en",
		"internal_signature": "",
		"customer_id": "test",
		"delivery_service": "meest",
		"shardkey": "9",
		"sm_id": 99,
		"date_created": "2021-11-26T06:22:19Z",
		"oof_shard": "1"
	}`

	so1 := `{
		"order_uid": "b563feb7b2b84b6test1",
		"track_number": "WBILMTESTTRACK1",
		"entry": "WBIL1",
		"delivery": {
		  "name": "Test Testov1",
		  "phone": "+97200000001",
		  "zip": "26398091",
		  "city": "Kiryat Mozkin1",
		  "address": "Ploshad Mira 151",
		  "region": "Kraiot1",
		  "email": "test@gmail.com1"
		},
		"payment": {
		  "transaction": "b563feb7b2b84b6test1",
		  "request_id": "",
		  "currency": "USD1",
		  "provider": "wbpay1",
		  "amount": 18171,
		  "payment_dt": 163790772,
		  "bank": "alpha1",
		  "delivery_cost": 15001,
		  "goods_total": 3171,
		  "custom_fee": 0
		},
		"items": [
		  {
			"chrt_id": 99349301,
			"track_number": "WBILMTESTTRACK1",
			"price": 4531,
			"rid": "ab4219087a764ae0btest1",
			"name": "Mascaras1",
			"sale": 301,
			"size": "0",
			"total_price": 3171,
			"nm_id": 23892121,
			"brand": "Vivienne Sabo1",
			"status": 2021
		  }
		],
		"locale": "en1",
		"internal_signature": "",
		"customer_id": "test1",
		"delivery_service": "meest1",
		"shardkey": "91",
		"sm_id": 99,
		"date_created": "2021-11-26T06:22:19Z",
		"oof_shard": "11"
	}`

	so2 := `{
		"order_uid": "b563feb7b2b84b6test2",
		"track_number": "WBILMTESTTRACK2",
		"entry": "WBIL2",
		"delivery": {
		  "name": "Test Testov2",
		  "phone": "+97200000002",
		  "zip": "26398092",
		  "city": "Kiryat Mozkin2",
		  "address": "Ploshad Mira 152",
		  "region": "Kraiot2",
		  "email": "test@gmail.com2"
		},
		"payment": {
		  "transaction": "b563feb7b2b84b6test2",
		  "request_id": "",
		  "currency": "USD2",
		  "provider": "wbpay2",
		  "amount": 18172,
		  "payment_dt": 163790771,
		  "bank": "alpha2",
		  "delivery_cost": 15002,
		  "goods_total": 3172,
		  "custom_fee": 0
		},
		"items": [
		  {
			"chrt_id": 99349302,
			"track_number": "WBILMTESTTRACK2",
			"price": 453,
			"rid": "ab4219087a764ae0btest2",
			"name": "Mascaras2",
			"sale": 30,
			"size": "0",
			"total_price": 3172,
			"nm_id": 23892122,
			"brand": "Vivienne Sabo2",
			"status": 202
		  }
		],
		"locale": "en2",
		"internal_signature": "",
		"customer_id": "test2",
		"delivery_service": "meest2",
		"shardkey": "9",
		"sm_id": 99,
		"date_created": "2021-11-26T06:22:19Z",
		"oof_shard": "1"
	}`

	var arr []string
	arr = append(arr, so, so1, so2)

	for i := 0; i < len(arr); i++ {
		sc.Publish("test.msg", []byte(arr[i]))
		//time.Sleep(2 * time.Second)
	}

}
