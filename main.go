package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/CalvoM/Unfurler"
	"github.com/go-redis/redis/v8"
)

type Response struct {
	data map[string]map[string]string
}

var ctx = context.Background()

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
		c, err := redis.ParseURL(os.Getenv("REDIS_URL"))
		if err != nil {
			log.Fatal(err)
		}
		c.Username = "" // Solve ERR wrong number of arguments of 'auth' error
		rdb := redis.NewClient(c)
		val, err := rdb.Get(ctx, u.Url).Result()
		if err == redis.Nil {
			uf := Unfurler.Unfurler{Url: u.Url}
			data := uf.Unfurl()
			jsonVal, _ := json.Marshal(data)
			_, err := rdb.Set(ctx, u.Url, jsonVal, 0).Result()
			if err != nil {
				log.Fatal(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonVal)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(val))
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method is not allowed")
	}

}

func main() {
	port := os.Getenv("PORT")
	fmt.Println("Running on", port)
	if port == "" {
		log.Fatal("PORT is not set")
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/unfurl/", unfurlHandler)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	_ = server.ListenAndServe()
}
