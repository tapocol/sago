package main

import (
  "github.com/craigjackson/sago"
  "github.com/craigjackson/sago_chat"
  "fmt"
)

func auth(r *sago.Request) {
  fmt.Printf("%i\n", r.Args)
  r.Session.Data["name"] = r.Args["name"]
  r.Session.Send(r.Id, "success", make(map[string]interface{}))
}

func main() {
  ac := sago.InitChannel("auth")
  go func() {
    for v := range ac {
      auth(v)
    }
  }()
  sago_chat.Init()
  sago.Run()
}

