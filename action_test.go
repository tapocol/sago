package sago

import (
  "testing"
)

var s = Session{}

func Test_AddAction(t *testing.T) {
  args := make(map[string]interface{})
  request := Request{Session: &s, Id: "1234", Args: &args}
  handler_called := false
  AddAction("foo", func(r *Request) {
    if r != &request {
      t.Error("Did not receive the expected request")
    }
    handler_called = true
  })
  if actions["foo"] == nil {
    t.Error("Did not initialize the name in actions.")
  }
  actions["foo"] <- &request
  if !handler_called {
    t.Error("The handler was never called")
  }
}

func Test_listen(t *testing.T) {
  c := make(chan *Request)
  args := make(map[string]interface{})
  request := Request{Session: &s, Id: "1234", Args: &args}
  handler_called := false
  go listen(c, func(r *Request) {
    if r != &request {
      t.Error("Did not receive the expected request")
    }
    handler_called = true
  })
  c <- &request
  if !handler_called {
    t.Error("The handler was never called")
  }
}

func Test_handle(t *testing.T) {
  actions["foo"] = make(chan *Request)
  id := "1234"
  args := make(map[string]interface{})
  go func() {
    request := <-actions["foo"]
    if (request.Session != &s) {
      t.Error("Did not receive the expected Session")
    }
    if (request.Id != id) {
      t.Error("Did not receive the expected Id")
    }
    if (request.Args != &args) {
      t.Errorf("Did not receive the expected Args %i != %i", request.Args, &args)
    }
  }()
  msg := message{Id: id, Action: "foo", Args: &args}
  handle(&s, msg)
}

