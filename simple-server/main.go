package main

import(
	"fmt"
	"log"
	"net/http"
)

func helloHandlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound);
		return;
	}
	if r.Method != "GET"{
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed);
        return;
	}
	fmt.Fprintf(w, "Hello, %s!")

}

func formHandlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	address:= r.FormValue("address")
	if name == "" {
		http.Error(w, "400 bad request", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Name, %s!", name)
	fmt.Fprintf(w, "Address, %s!", address)
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"));
	http.Handle("/", fileServer);
	http.HandleFunc("/form",formHandlerFunc);
	http.HandleFunc("/hello",helloHandlerFunc);

	fmt.Printf("Starting server at port 8000");
	if err := http.ListenAndServe(":8000",nil); err != nil{
		log.Fatal(err);
	}
}