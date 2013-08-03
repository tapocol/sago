package sago

import (
)

type Request struct {
  Session *Session
  Id string
  Args map[string]interface{}
}

var actions = make(map[string]chan *Request)

func AddAction(name string, handler func(*Request)) {
  actions[name] = make(chan *Request)
  go func() {
    for {
      handler(<-actions[name])
    }
  }()
}

func handle(s *Session, msg message) {
  actions[msg.Action] <- &Request{Session: s, Id: msg.Id, Args: msg.Args}
}

