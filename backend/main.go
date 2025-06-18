package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "golang.org/x/time/rate"
    "github.com/gorilla/handlers"
)

var limiter = rate.NewLimiter(1, 3)

func secureHandler(w http.ResponseWriter, r *http.Request) {
    if !limiter.Allow() {
        http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, `{"message": "Hello from Go backend!"}`)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/api", secureHandler)

    handler := handlers.CORS(
        handlers.AllowedOrigins([]string{"*"}),
        handlers.AllowedMethods([]string{"GET"}),
    )(mux)

    srv := &http.Server{
        Addr:              ":8080", // 🔧 listen on port 8080
        Handler:           handler,
        ReadHeaderTimeout: 5 * time.Second,
    }

    log.Println("✅ Backend running on :8080")
    log.Fatal(srv.ListenAndServe())
}

