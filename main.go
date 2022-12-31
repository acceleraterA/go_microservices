package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/acceleraterA/go_microservices/details"
	"github.com/gorilla/mux"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("checking application health")
	response := map[string]string{
		"status":    "UP",
		"timestamp": time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)
}
func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("serving the home page")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Application is up and running")
}
func detailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching the details")
	hostname, err := details.Gethostname()

	if err != nil {
		panic(err)
	}
	IP, _ := details.GetIP()
	fmt.Println(hostname, IP)
	//fmt.Fprintf(w, hostname, IP)
	response := map[string]string{
		"hostname": hostname,
		"IP":       IP.String(),
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/details", detailsHandler)
	log.Println("Server has started!!!")
	log.Fatal(http.ListenAndServe(":80", r))

}

/*package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, you've requested: %s with token: %s \n", r.URL.Path, r.URL.Query().Get("token"))
}

func main() {

	http.HandleFunc("/", rootHandler)
	//use http.FileServer and point it to a url path here is dir(static).
	fs := http.FileServer(http.Dir("static/"))
	fmt.Println(fs)
	//point a url path to fs
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Println("Web server has started")
	http.ListenAndServe(":80", nil)
}
/*
package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s with token: %s \n", r.URL.Path, r.URL.Query().Get("token"))
}

func main() {
	http.HandleFunc("/", rootHandler)
	//use http.FileServer and point it to a url path here is dir(static).
	fs := http.FileServer(http.Dir("static/"))
	fmt.Println(fs)
	//point a url path to fs
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Println("Web server has started")
	http.ListenAndServe(":80", nil)
}

package main

import (
	"fmt"
	"runtime"
	"unsafe"

	//unsafe should be used with care
	geo "github.com/acceleraterA/go_microservices/geometry"

	"rsc.io/quote"
)

func rectProps(length, width float64) (area, perimeter float64) {
	//var area = length * width
	area = length * width
	//var perimeter = (width + length) * 2
	perimeter = (width + length) * 2
	return
}

func main() {
	x := 10
	var y, z = 2, 3
	//var name string
	isworking := false
	name := "Devops"
	fmt.Println("Hello world!")
	//random quote
	fmt.Println(quote.Go())
	fmt.Println(x, y, z, name, isworking)
	//the type of variable can be printed using %T
	//printIn automatically change line,
	//printf need to add \n to change line
	fmt.Printf("Type of name is %T, size of name is %d \n", name, unsafe.Sizeof(name))
	a, b := rectProps(2, 3)
	fmt.Printf("Area is %f and perimeter is %f \n", a, b)
	daysOfMonth := map[string]int{"Jan": 31, "Feb": 28}
	fmt.Println(daysOfMonth)
	area := geo.Area(1, 2)
	diag := geo.Diagnal(1, 2)
	fmt.Println(area, diag)
	fmt.Printf("OS: %s\nArch: %s\n", runtime.GOOS, runtime.GOARCH)
*/
