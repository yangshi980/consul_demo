// +build ignore

package main

import (
    "fmt"
    "log"
    "net/http"
    "strings"
    "consul_demo/demo"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
  query:=r.URL.Query()
  vec:=make([]string,0)
  for k,v := range query{
    ss:=fmt.Sprintf("key:%s,value:%s" ,k,v)
    //fmt.Println(ss)
    vec=append(vec, ss)
  }
  rrr:=strings.Join(vec,"&&")
  //fmt.Println(string(rrr))
  fmt.Fprintf(w, "Hi there, I love %s! para is %s", r.URL.Path[1:],rrr)
}

func main() {
    http.HandleFunc("/", handler)
    go func (){
        time.Sleep(2 * 1000 * time.Millisecond)
        client,err :=demo.Reg("127.0.0.1","demo_server","demo_server_id001",8080)
      if err!=nil{
        panic(err)
      }
      demo.Watch(client,"demo/testing")
    }()
    log.Fatal(http.ListenAndServe(":8080", nil))
}
