package main

import (
  "net/http"
  "fmt"
  "net/http/httputil"
//  "bytes"
)

func main() {
  http.HandleFunc("/", RequestHandler)
  http.ListenAndServe(":8080", nil)
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
  dump, _ := httputil.DumpRequest(r, true)
  w.Write(dump)
  
  fmt.Printf("  ---[ Request from %s ]---\n", r.RemoteAddr)
 // n := bytes.IndexByte(dump, 0)
  fmt.Println(string(dump[:]))

}
