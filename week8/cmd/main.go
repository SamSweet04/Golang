package main

import (
	"net/http"
	"fmt"
	// m "github.com/SamSweet04/Golang/week8/models"
	f "github.com/SamSweet04/Golang/week8/pkg"
)

func main(){
	result, err := db.Exec("insert into users (name, surname, login, password) values ('Amina', 'Amangeldi', 'login', 'passwprd')")
	if err != nil{
		panic(err)
	}
	u := m.User{"name", "surname", "lap", "djk"}
	f.Register(u)
	f.AddItem(m.Item{"Samsung S22", 450000, 0})
	f.Rate(m.Rating{"Samsung S22", 5})
	f.SearchItemByName("Samsung S22")
	Authorize(Authorization{"ff", "d"})
	AddItem(Item{"Iphone 12", 15000, 0})
	SearchItemByName("Iphone 12")
	Rate(Rating{"Iphone 12", 4})
	Rate(Rating{"Iphone 12", 3})
	Rate(Rating{"Iphone 12", 2})
	f.SearchItemByName("Iphone 12")
	FilterByPrice(10000, 200000)
	f.FilterByRating(2, 4)
	http.HandleFunc("/register", f.Register)
	http.HandleFunc("/authorize", f.Authorize)
	http.HandleFunc("/additem", f.AddItem)
	http.HandleFunc("/byname", f.SearchItemByName)
	http.HandleFunc("/rating", f.FilterByRating)
	http.HandleFunc("/price", f.FilterByPrice)
	http.HandleFunc("/giverating", f.Rate)
	fmt.Println("Server: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}