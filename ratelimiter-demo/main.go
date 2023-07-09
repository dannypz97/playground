package main

import (
	"fmt"
	"io"
	"net/http"

	limiter "github.com/dannypz97/ratelimiter/limiter"
)

var rl *limiter.RateLimiter

func init() {
	rl = limiter.NewRateLimiter(25, 10)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, fmt.Sprintf("%v", rl.IsAllowed()))
}

func main() {
	http.HandleFunc("/", getRoot)
	http.ListenAndServe(":8080", nil)
}
