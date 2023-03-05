package main
import ("fmt"
"net/http")

type ViewData struct {
	Title string
	Message string
}
func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	
		data := ViewData{
		Title: "World Cup",
		Message: "FIFA will never regret it",}
	tmpl := template.Must(template.New( "data").Parse("<div>
	<h1>{{ .Title}}</h1>
	 <p>{{ .Message}}</p>
	</div>" ))
	tmpl.Execute(w, data)
		})
	fmt.Println( "Server is listening.…")
	err := http.ListenAndServe( ":8181", nil)
	if err != nil{
	return
	}
}