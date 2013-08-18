package main

import (
  "github.com/craigjackson/sago"
  "github.com/craigjackson/sago_chat"
)

type AuthArgs struct {
  Name string `json:"name"`
}

func auth(r *sago.Request) {
  args := AuthArgs{}
  r.Unmarshal(&args)
  r.Session.Data["name"] = args.Name
  r.Session.Send(r.Id, "success", make(map[string]interface{}))
}

func main() {
  sago.AddAction("auth", auth)
  sago_chat.Init()
  sago.Run()
}

