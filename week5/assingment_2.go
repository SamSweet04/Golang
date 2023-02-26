package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
	"strings"
)
// User represents a user in the system
type User struct {
	Username string
	Password string
}

type Registration struct {
	Name     string
	Surname  string
	Age      int
	Login    string
	Password string
}

type Authorization struct {
	Login    string
	Password string
}

type Item struct {
	Name       string
	Price      float64
	Rating     float64
	RatingList []float64
}

type Database struct {
	Logins []Registration
	Items  []Item
	Users  []User
}

func (d *Database) LoadUsers() error{
	file, err := os.Open("users.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}
		user := &User{
			Username: parts[0],
			Password: parts[1],
		}
		d.Users = append(d.Users, *user)
	}
	return scanner.Err()
}

func (d *Database) Register(r Registration) *Registration {
	for i := 0; i < len(d.Logins); i++ {
		if d.Logins[i].Login == r.Login && d.Logins[i].Password == r.Password {
			// return fmt.Errorf("username %s %s already exists", r.Name, r.Surname)
			fmt.Printf("User %s %s already exists!\n", r.Name, r.Surname)
			return nil
		}
	}
	reg := &Registration{r.Name, r.Surname, r.Age, r.Login, r.Password}
	d.Logins = append(d.Logins, Registration{r.Name, r.Surname, r.Age, r.Login, r.Password})
	fmt.Println("You registred!")
	return reg
}

func (a *Authorization) SignIn(d Database) string {
	for i := 0; i < len(d.Logins); i++ {
		if d.Logins[i].Login == a.Login && d.Logins[i].Password == a.Password {
			return "You entered system!"
		}
	}
	return "No authorized!!!"
}

func (d *Database) AddItem(name string, price float64) *Item {
	item := &Item{name, price, 0, nil}
	d.Items = append(d.Items, Item{name, price, 0, nil})
	return item
}

func (d *Database) SearchItem(name string) {
	for i := 0; i < len(d.Items); i++ {
		if strings.Contains(strings.ToUpper(d.Items[i].Name), strings.ToUpper(name)) {
			fmt.Println(d.Items[i])
			return
		} else {
			fmt.Printf("Did not find an item like %s", name)
		}
	}
}

func (d *Database) FilteringItems(price1, price2, rating1, rating2 float64) {
	for i := 0; i < len(d.Items); i++ {
		if d.Items[i].Price >= price1 && d.Items[i].Price <= price2 && d.Items[i].Rating >= rating1 && d.Items[i].Rating <= rating2 {
			fmt.Println("We found item that you searched!")
			fmt.Println(d.Items[i])
			return
		
		} 
	}
	fmt.Println("No such item with these price and rating!!!")
}

func Rate(d Database, a Authorization, itemName string, rating float64) {
	if a.SignIn(d) == "You entered system!" {
		var sum float64
		for i := 0; i < len(d.Items); i++ {
			if d.Items[i].Name == itemName {
				d.Items[i].RatingList = append(d.Items[i].RatingList, rating)
				fmt.Println("You successfully rated an item!")
				for j := 0; j < len(d.Items[i].RatingList); j++ {
					sum += d.Items[i].RatingList[j]
				}
				d.Items[i].Rating = sum / float64(len(d.Items[i].RatingList))
			}else{
				continue
			}
		}
		
	} else {
		fmt.Println("UNKNOWN!!!")
	}
}
func main() {
	// var ar [] a.Registration
	d := Database{}
	d.AddItem("IPhone", 500000)
	d.AddItem("Nokia", 100000)
	r := Registration{Name: "Saule", Surname: "Arystanbek", Age: 18, Login: "l", Password: "p"}
	// r1 := Registration{Name: "Jin", Password: "Kim", Age: 24, Login: "l1", Password: "p1"}
	r1 := Registration{"Jin", "Kim", 24, "l1", "p1"}
	d.Register(r)
	d.Register(r1)
	// fmt.Println(d)
	// r1.Register(d)
	au := Authorization{Login: "l", Password: "p"}
	response := au.SignIn(d)
	fmt.Println(response)
	au1 := Authorization{Login: "l1", Password: "p1"}
	response1 := au1.SignIn(d)
	fmt.Println(response1)
	d.SearchItem("IPhone")
	Rate(d, au, "IPhone", 6)
	fmt.Println(d.Items)
	Rate(d, au1, "IPhone", 7)
	Rate(d, au, "Nokia", 10)
	d.FilteringItems(5000, 1000000, 4, 8)
	fmt.Println(d.Items)
}
