package sago

import (
  "encoding/json"
)

type Request struct {
  Session *Session
  Id string
  Args string
}

func (r *Request) Unmarshal(v interface{}) {
  json.Unmarshal([]byte(r.Args), v)
}

var actions = make(map[string]chan *Request)

func AddAction(name string, handler func(*Request)) {
  actions[name] = make(chan *Request)
  go listen(actions[name], handler)
}

func listen(c chan *Request, handler func(*Request)) {
  for {
    handler(<-c)
  }
}

func handle(s *Session, msg message) {
  actions[msg.Action] <- &Request{Session: s, Id: msg.Id, Args: msg.Args}
}

