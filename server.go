package main
import "net/http"
import "os"
import "io/ioutil"
import "time"
import "fmt"

var startedAt = time.Now()

func main() {
	http.HandleFunc("/health", Health)
	http.HandleFunc("/configmap",ConfigMap)
	http.HandleFunc("/",Hello)
	http.HandleFunc("/secret", Secret)
	http.ListenAndServe(":80",nil)
	
}

func Hello(w http.ResponseWriter, r *http.Request) {

	name := os.Getenv("NAME")
	w.Write([]byte("<h1>Hello, " + name + "</h1>"))
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("/go/myfamily/family.txt")
	if err != nil {
		w.Write([]byte("<h1>Error</h1>"))
	}
	w.Write([]byte("<h1>Hello Family: " + string(data) + "</h1>"))
}

func Secret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	w.Write([]byte("<h1>User/Password: " + user + " / " + password + "</h1>"))
}

func Health(w http.ResponseWriter, r *http.Request) {
	duration := time.Since(startedAt)

	if duration.Seconds() > 25 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}
}