package main

import (
  "github.com/craigjackson/sago"
  "github.com/craigjackson/sago/session"
  "fmt"
)

func auth(r *session.Request) {
  fmt.Printf("%i\n", r.Args)
  r.Session.Data["name"] = r.Args["name"]
  r.Session.Send(r.Id, "success", make(map[string]interface{}))
}

func chat(r *session.Request) {
  if r.Session.Data["name"] == nil {
    res := make(map[string]interface{})
    res["message"] = "Need to auth"
    r.Session.Send(r.Id, "fail", res)
    return
  }

  if r.Args["message"] == nil {
    res := make(map[string]interface{})
    res["message"] = "Need to include the message"
    r.Session.Send(r.Id, "fail", res)
    return
  }

  res := make(map[string]interface{})
  res["name"] = r.Session.Data["name"]
  res["message"] = r.Args["message"]
  session.SendAll("chat", res)
}

func main() {
  ac := session.InitChannel("auth")
  go func() {
    for v := range ac {
      auth(v)
    }
  }()
  cc := session.InitChannel("chat")
  go func() {
    for v := range cc {
      chat(v)
    }
  }()
  sago.Run()
}

