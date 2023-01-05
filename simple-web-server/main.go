package main
import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":		
		 http.ServeFile(w, r, "./static/form.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "POST request successful\n")
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", address)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
	}

	fmt.Fprintf(w, "hello!")
}

func main() {
	// create file server handler
	fs := http.FileServer(http.Dir("./static"))

	// handle '/' route
	http.Handle("/", fs)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080")
/**
*	nil is the zero value for pointers, interfaces, maps, slices, 
*	channels and function types, representing an uninitialized value.
*/
	// Run the web server.
	// func ListenAndServe(addr string, handler Handler) error
	if err := http.ListenAndServe(":8080", nil); err != nil {
		// Fatal is equivalent to Print() followed by a call to os.Exit(1).
		log.Fatal(err)
	}
}