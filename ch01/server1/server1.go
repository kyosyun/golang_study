package server1

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	// アクセスをhandlerに飛ばす
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}