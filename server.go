package main
import "net/http"
import "os"
import "io/ioutil"

func main() {
	http.HandleFunc("/configmap",ConfigMap)
	http.HandleFunc("/",Hello)
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