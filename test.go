package main

import (
       "fmt"
       "net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
       fmt.Fprintf(w, "<h1>Hello !</h1>")
}


func main() {
       http.HandleFunc("/", index)
       http.ListenAndServe("192.168.203.80:5000", nil)
       fmt.Println("伺服器已啟動...")
}
//1231231