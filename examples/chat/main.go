package main

import (
  "github.com/craigjackson/sago"
  "github.com/craigjackson/sago/session"
  "fmt"
)

func auth(s *session.Session, id string, data map[string]interface{}) {
  fmt.Printf("%i\n", data)
  s.Data["name"] = data["name"]
  s.Send(id, "success", make(map[string]interface{}))
}

func chat(s *session.Session, id string, data map[string]interface{}) {
  if s.Data["name"] == nil {
    res := make(map[string]interface{})
    res["message"] = "Need to auth"
    s.Send(id, "fail", res)
    return
  }

  if data["message"] == nil {
    res := make(map[string]interface{})
    res["message"] = "Need to include the message"
    s.Send(id, "fail", res)
    return
  }

  res := make(map[string]interface{})
  res["name"] = s.Data["name"]
  res["message"] = data["message"]
  session.SendAll("chat", res)
}

func main() {
  session.AddAction("auth", auth)
  session.AddAction("chat", chat)
  sago.Run()
}

