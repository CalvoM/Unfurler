package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/CalvoM/Unfurler"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	str := `<html>
	<head><title>Go! GO! Go!</title></head>
	<body>
		<h2 style="color:blue;font-family:Arial;letter-spacing:0.1em;text-align:center;">Let us Unfurl...Gentlemen shall we?</h2>
	</body>
	</html>
	`
	w.Write([]byte(str))
}

func unfurlHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method is not allowed")
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var u Unfurler.Unfurler
		err := decoder.Decode(&u)
		if err != nil {
			log.Fatal(err)
		}
		uf := Unfurler.Unfurler{Url: u.Url}
		data := uf.Unfurl()
		jsonVal, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonVal)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method is not allowed")
	}

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/unfurl/", unfurlHandler)
	server := &http.Server{
		Addr:"0.0.0.0:5000",
		Handler: mux,
	}
	_ = server.ListenAndServe()
}
