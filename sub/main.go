package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"stanService/model"
	"stanService/sub/dbCon"
	"sync"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
)

const (
	clusterName = "test-cluster"
	clientName  = "sub1"
	channel     = "orders"
)

var Cache = make(map[string]model.Order)

//Получаем Order и кладем в map
func getOrderFromDb(db *sql.DB, id string) model.Order {
	var (
		items []model.Item
		o     model.Order
		d     model.Delivery
		i     model.Item
		p     model.Payment
	)

	result, err := db.Query(fmt.Sprintf("SELECT * FROM orders where Order_uid ='%s'", id))
	if err != nil {
		panic(err)
	}
	for result.Next() {
		err := result.Scan(&o.Order_uid, &o.Track_number, &o.Entry, &o.Locale, &o.Internal_signature, &o.Customer_id, &o.Delivery_service, &o.Shardkey, &o.Sm_id, &o.Date_created, &o.Oof_shard)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	result2, err := db.Query(fmt.Sprintf("SELECT * FROM delivery where order_id ='%s'", id))
	if err != nil {
		panic(err)
	}
	for result2.Next() {
		err := result2.Scan(&d.Name, &d.Phone, &d.Zip, &d.City, &d.Address, &d.Region, &d.Email, &d.Id)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	result3, err := db.Query(fmt.Sprintf("SELECT chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status FROM items where order_id ='%s'", id))
	if err != nil {
		panic(err)
	}
	for result3.Next() {
		err := result3.Scan(&i.Chrt_id, &i.Track_number, &i.Price, &i.Rid, &i.Name, &i.Sale, &i.Size, &i.Total_price, &i.Nm_id, &i.Brand, &i.Status)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	items = append(items, i)
	result4, err := db.Query(fmt.Sprintf("SELECT transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee FROM payment where order_id ='%s'", id))
	if err != nil {
		panic(err)
	}
	for result4.Next() {
		err := result4.Scan(&p.Transaction, &p.Request_id, &p.Currency, &p.Provider, &p.Amount, &p.Payment_dt, &p.Bank, &p.Delivery_cost, &p.Goods_total, &p.Custom_fee)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	o.Delivery = d
	o.Payment = p
	o.Items = items
	//Если order_uid не упустой, кладем в кэш
	if o.Order_uid != "" {
		Cache[o.Order_uid] = o
		fmt.Println("Положили в кеш:", o)
	}
	return o
}

// Выбираем откуда выдавать ордер(кэш или БД)
func selectOrder(id string, db *sql.DB) model.Order {
	//Если нашли такой ключ с таким id в кэше, отдаем знаечение из кэша
	if val, found := Cache[id]; found {
		return val
	}
	//Если значение в кэше нет, то идем в базу
	gg := getOrderFromDb(db, id)
	return gg
}

func main() {
	sc, _ := stan.Connect(clusterName, clientName)
	defer sc.Close()
	dbCon.InsertValue(sc)
	time.Sleep(5 * time.Second)
	handleFunc()

	w := sync.WaitGroup{}
	w.Add(1)
	w.Wait()

}
func handleFunc() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/{id}", getOrderById).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", rtr))
}

func getOrderById(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./httpServ/temp/template.html")
	if err != nil {
		fmt.Println(err)
	}
	vars := mux.Vars(r)
	db := dbCon.DbConnection()
	defer db.Close()
	id := vars["id"]
	gg := selectOrder(id, db)
	if gg.Order_uid == id {
		t.Execute(w, gg)
		fmt.Println("Вывели gg", gg)
	} else {
		fmt.Fprintf(w, "not order")
	}

}
