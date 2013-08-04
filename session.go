package sago

import (
  "fmt"
  "code.google.com/p/go.net/websocket"
)

type message struct {
  Id string `json:"id"`
  Action string `json:"action"`
  Args *map[string]interface{} `json:"args"`
}

type Session struct {
  ws *websocket.Conn
  Data map[string]interface{}
}

var live_sessions = map[*Session]bool{}

func addLiveSession(s *Session) {
  live_sessions[s] = true
}

func removeLiveSession(s *Session) {
  delete(live_sessions, s)
}

func (s Session) Start() {
  fmt.Println("Connected:", s.ws)
  addLiveSession(&s)
  s.listen()
  removeLiveSession(&s)
  fmt.Println("Disconnected:", s.ws)
}

func (s Session) listen() {
  for {
    var msg message
    err := websocket.JSON.Receive(s.ws, &msg)
    if err != nil {
      fmt.Println(err)
      break
    }
    fmt.Printf("Session: %i, Message: %i\n", s, msg)
    handle(&s, msg)
  }
}

func (s Session) Send(response_id string, name string, args map[string]interface{}) error {
  msg := message { Id: response_id, Action: name, Args: &args }
  return websocket.JSON.Send(s.ws, msg)
}

func SendAll(name string, args map[string]interface{}) {
  SendAllExcept(name, args, make([]*Session, 0))
}

func SendAllExcept(name string, args map[string]interface{}, exceptions []*Session) {
  var excepted bool
  for k := range live_sessions {
    excepted = false
    for _, v := range exceptions {
      if v == k {
        excepted = true
        break
      }
    }
    if !excepted {
      k.Send("", name, args)
    }
  }
}

