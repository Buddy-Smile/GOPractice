package main
 
import (
    "encoding/json"
    "math/rand"
    "net/http"
    "time"
)
 
var quotes = []string{
    "Believe you can and you're halfway there.",
    "Do not watch the clock. Do what it does. Keep going.",
    "Everything you can imagine is real.",
    "The future belongs to those who believe in the beauty of their dreams.",
}
 
func getQuote(w http.ResponseWriter, r *http.Request) {
    rand.Seed(time.Now().UnixNano())
    quote := quotes[rand.Intn(len(quotes))]
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"quote": quote})
}
 
func main() {
    http.HandleFunc("/quote", getQuote)
    http.ListenAndServe(":8080", nil)
}
