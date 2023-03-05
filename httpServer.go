package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w, r, "static/user.html")
	})
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request){
		name := r.FormValue("username")
		age := r.FormValue("userage")
		fmt.Fprintf(w, "Name %s; Age %s",name,age)
	})
	// fs := http.FileServer(http.Dir("static"))
	// http.Handle("/",fs)

	// http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request){
	// 	fmt.Fprint(w, "About page")
	// })

	// http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request){
	// 	fmt.Fprint(w, "Contact page")
	// })

	fmt.Println("Server is listening…")
	// httpHandler := http.FileServer(http.Dir("static"))
	err := http.ListenAndServe(":8181", nil)
	if err != nil {
		return
	}
}

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, format: "This is a homepage")
// 	 fmt.Println("Endpoint hit: homePage")
// }
// func handleRequests() {
// 	http.HandleFunc(@v"/", homePage)
// 	fmt.Println( "Server is listening…" )
// 	err := http.ListenAndServe( addr: ":8181", handler: nil)
// 	if err != nil {
// 	fmt.Println(err. Error())
// 	}
// }
// func main() {
// handleRequests()
// // }
// func getRoot(w http.ResponseWriter, r *http.Request) {
//     fmt.Printf(  "got / request\n")
// 	 io.WriteString(w, "This is my website!\n")
// }
// func getHello(w http.ResponseWriter, r *http.Request) {
// 	 fmt.Printf(  "got /hello request\n")
// 	  io.WriteString(w, "Hello, HTTP!\n")
// }
// func main() {
//     http.HandleFunc("/", getRoot)
// 	 http.HandleFunc("/hello", getHello)
// 	 fmt.Println("Server is listening… ")
// 	 err := http.ListenAndServe( "18181",  nil)
// 	 if err != nil
// 	}
