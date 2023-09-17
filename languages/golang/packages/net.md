`net/http`
```
resp, err := http.Get(url)

http.HandleFunc("/", handler)
func handler( w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "URL path = %q\n", r.URL.Path)
}
```