package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
	"html/template"
	"log"
	"net/http"
	"stanService/model"
	"stanService/sub/dbCon"
	"sync"
	"time"
)

const (
	clusterName = "test-cluster"
	clientName  = "sub1"
	channel     = "orders"
)

//Получаем Order и кладем в map
func getOrder(db *sql.DB, id string) model.Order {

	var items []model.Item
	var o model.Order
	var d model.Delivery
	var i model.Item
	var p model.Payment
	//m := make(map[string]model.Order)

	result, err := db.Query(fmt.Sprintf("SELECT * FROM orders where Order_uid ='%s'", id))
	if err != nil {
		panic(err)
	}
	for result.Next() {
		//o := Order{}
		err := result.Scan(&o.Order_uid, &o.Track_number, &o.Entry, &o.Locale, &o.Internal_signature, &o.Customer_id, &o.Delivery_service, &o.Shardkey, &o.Sm_id, &o.Date_created, &o.Oof_shard)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	result2, err := db.Query("SELECT * FROM delivery")
	if err != nil {
		panic(err)
	}
	for result2.Next() {
		err := result2.Scan(&d.Name, &d.Phone, &d.Zip, &d.City, &d.Address, &d.Region, &d.Email, &d.Id)
		if err != nil {
			panic(err)
			continue
		}
	}
	result3, err := db.Query("SELECT * FROM item")
	if err != nil {
		panic(err)
	}
	for result3.Next() {
		err := result3.Scan(&i.Chrt_id, &i.Track_number, &i.Price, &i.Rid, &i.Name, &i.Sale, &i.Size, &i.Total_price, &i.Nm_id, &i.Brand, &i.Status)
		if err != nil {
			panic(err)
			continue
		}
	}
	items = append(items, i)
	result4, err := db.Query(fmt.Sprintf("SELECT * FROM payment where transaction ='%s'", id))
	if err != nil {
		panic(err)
	}
	for result4.Next() {
		err := result4.Scan(&p.Transaction, &p.Request_id, &p.Currency, &p.Provider, &p.Amount, &p.Payment_dt, &p.Bank, &p.Delivery_cost, &p.Goods_total, &p.Custom_fee)
		if err != nil {
			panic(err)
			continue
		}
	}
	o.Delivery = d
	o.Payment = p
	o.Items = items

	//fmt.Println(i)
	//fmt.Println(p)
	//fmt.Println(d)
	//fmt.Println(o)
	//if _, found := m[id]; found {
	//	return m[id]
	//	fmt.Println("Значение из мапы ", m[id])
	//}
	//m[o.Order_uid] = o

	fmt.Println("значение из бд", o)
	return o
}

func main() {
	sc, _ := stan.Connect(clusterName, clientName)
	defer sc.Close()
	//insertValue(sc)
	dbCon.InsertValue(sc)
	time.Sleep(5 * time.Second)
	handleFunc()

	w := sync.WaitGroup{}
	w.Add(1)
	w.Wait()

}
func handleFunc() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/{id}", handleTest).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", rtr))
	//http.ListenAndServe(":8081", nil)
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./httpServ/temp/template.html")
	if err != nil {
		fmt.Println(err)
	}

	vars := mux.Vars(r)

	db := dbCon.DbConnection()
	defer db.Close()

	id := vars["id"]
	gg := getOrder(db, id)
	//fmt.Fprintf(w, "ID:", vars[id], gg)
	t.Execute(w, gg)

}
