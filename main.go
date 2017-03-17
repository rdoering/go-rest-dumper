package main

import (
  "net/http"
  "fmt"
  "net/http/httputil"
)

func main() {
  http.HandleFunc("/", RequestHandler)
  http.ListenAndServe(":8080", nil)
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
  dump, _ := httputil.DumpRequest(r, true)
  //w.Write([]byte("hello"))
  w.Write(dump)
  
  fmt.Println("  ---[ Request from  ]---")
  fmt.Println("${request.method} /${request.url} HTTP/${request.protocolVersion}");
  fmt.Println("Host: ");
  /*request.headers.forEach((key, value){
    print("$key: $value");
  });*/
  /*if(request.contentLength != null) {
    request.readAsString().then(print);
  }*/

}
