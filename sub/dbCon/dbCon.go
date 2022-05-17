package dbCon

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"stanService/model"
)

// InsertValue Подписываемся на канал
func InsertValue(sc stan.Conn) {
	sc.Subscribe("test.msg", MsgHandler) //stan.DurableName("my-durable") для того чтоб не терять данные
}

func DbConnection() *sql.DB {
	connDb := "user=me password=password dbname=medb sslmode=disable"
	db, err := sql.Open("postgres", connDb)
	if err != nil {
		panic(err)
	}
	return db
}

// MsgHandler Получаем данные из сообщения, парсим и кладем в базу.
func MsgHandler(m *stan.Msg) {
	db := DbConnection()
	defer db.Close()

	var order model.Order

	err := json.Unmarshal(m.Data, &order)
	if err != nil {
		fmt.Println(err)
	}
	_, err = db.Exec(fmt.Sprintf("insert into orders (order_uid, track_number, entry, locate, internal_signature, customer_id, delivery_service, shardkey, date_created, oof_shard) values ('%s','%s', '%s', '%s','%s','%s','%s','%s','%s','%s')",
		order.Order_uid, order.Track_number, order.Entry, order.Locale, order.Internal_signature,
		order.Customer_id, order.Delivery_service, order.Shardkey, order.Date_created, order.Oof_shard))
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(fmt.Sprintf("insert into delivery values ('%s','%s','%s','%s','%s','%s','%s')",
		order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City,
		order.Delivery.Address, order.Delivery.Region, order.Delivery.Email))
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(fmt.Sprintf("insert into item (track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) values ('%s', '%d', '%s', '%s', '%d', '%s', '%d','%d','%s','%d')",
		order.Items[0].Track_number, order.Items[0].Price, order.Items[0].Rid, order.Items[0].Name, order.Items[0].Sale,
		order.Items[0].Size, order.Items[0].Total_price, order.Items[0].Nm_id, order.Items[0].Brand, order.Items[0].Status))
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(fmt.Sprintf("insert into payment values ('%s','%s','%s','%s','%d','%d','%s','%d','%d','%d')",
		order.Payment.Transaction, order.Payment.Request_id, order.Payment.Currency, order.Payment.Provider, order.Payment.Amount,
		order.Payment.Payment_dt, order.Payment.Bank, order.Payment.Delivery_cost, order.Payment.Goods_total, order.Payment.Custom_fee))
	if err != nil {
		panic(err)
	}
}
