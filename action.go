package sago

import (
)

type Action struct {
  Channel chan *Request
}

type Request struct {
  Session *Session
  Id string
  Args map[string]interface{}
}

var actions = make(map[string]Action)

func AddAction(name string) Action {
  actions[name] = Action{Channel: make(chan *Request)}
  return actions[name]
}

func handle(s *Session, msg message) {
  actions[msg.Action].Channel <- &Request{Session: s, Id: msg.Id, Args: msg.Args}
}

