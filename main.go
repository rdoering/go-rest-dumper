package main

import (
  "net/http"
  "fmt"
)

func main() {

  http.HandleFunc("/", RequestHandler)
  http.ListenAndServe(":8080", nil)
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("  ---[ Request from  ]---")
  fmt.Println("${request.method} /${request.url} HTTP/${request.protocolVersion}");
  fmt.Println("Host: ");
  /*request.headers.forEach((key, value){
    print("$key: $value");
  });*/
  /*if(request.contentLength != null) {
    request.readAsString().then(print);
  }*/

    w.Write([]byte("hello!"))
}
