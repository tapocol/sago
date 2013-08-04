package main

import (
  "github.com/craigjackson/sago"
  "github.com/craigjackson/sago_chat"
)

func auth(r *sago.Request) {
  args := *r.Args
  r.Session.Data["name"] = args["name"]
  r.Session.Send(r.Id, "success", make(map[string]interface{}))
}

func main() {
  sago.AddAction("auth", auth)
  sago_chat.Init()
  sago.Run()
}

