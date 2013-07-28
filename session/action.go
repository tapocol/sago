package session

import (
)

var handlers = make(map[string]func(s *Session, id string, args map[string]interface{}))

func AddAction(name string, handler func(s *Session, id string, args map[string]interface{})) {
  handlers[name] = handler
}

func execute(name string, s *Session, id string, args map[string]interface{}) bool {
  if handlers[name] == nil {
    return false
  }
  handlers[name](s, id, args)
  return true
}

