package session

import (
  "fmt"
)

type Request struct {
  Session *Session
  Id string
  Args map[string]interface{}
}

var channels = make(map[string]chan *Request)

func InitChannel(name string) chan *Request {
  c := make(chan *Request)
  channels[name] = c
  return c
}

func handle(s *Session, msg message) {
  if channels[msg.Action] != nil {
    channels[msg.Action] <- &Request{Session: s, Id: msg.Id, Args: msg.Args}
  } else {
    fmt.Println("Action", msg.Action, "has no channel")
  }
}

