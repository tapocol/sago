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
  sago.AddAction("auth", auth)
  sago_chat.Init()
  sago.Run()
}

