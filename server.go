package main
import "net/http"
import "os"

func main() {
	http.HandleFunc("/",Hello)
	http.ListenAndServe(":80",nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {

	name := os.Getenv("NAME")
	w.Write([]byte("<h1>Hello, " + name + "</h1>"))
}