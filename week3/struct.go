package main
import "fmt"

type person struct{
	name string
	age int
}
func main(){
	tom := person {name : "Tom", age: 22}
	var tomPointer *person = &tom
	tomPointer.age = 29
	fmt.Println(tom.age)
	(*tomPointer).age = 32
	fmt.Println(tom.age)
}

/*
tom := person{name: "Tom", age: 22}
var agePointer *int = &tom.age
*agePointer = 35
fmt.Println(tom.age)
*/

/*
var tom person = person{"Tom", 23}
bob := person{ name : " Bob", age : 31}
*/

// Вложенные структуры
/*
type contact struct{
	email string
	phone string
}
type person struct{
	name string
	age int
	contactinfo contact
}
func main(){
	var tom = person{
		name: "Tom",
		age: 24,
		contactInfo: contact{
			email: "tom@.gmail.com"
			phone: "+7076665404"
		}, 
	}
	tom.contactInfo.email = "Saule@gmail.com"
	fmt.Println(tom.contactInfo.email)
	fmt.Println(tom.contactInfo.phone)
}
*/

// Хранение ссылки на структуру того же типа
/*
type node struct{
	value int
	next node --- it is not correct
} 
type node struct{
	value int 
	next *node
}
func printNodeValue(n *node){
	fmt.Println(n.value)
	if(n.next != nil){
		printNodeValue(n.next)
	}
}
func main(){
	first := node{value: 4}
	second := node{value : 5}
	third := node{value : 6}

	first.next = &second
	second.next = &third

	var current *node = &first
	for current != nil{
		fmt.Println(current.value)
		current = current.next
	}
}
*/