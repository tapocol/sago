package sago

import (
  "testing"
)

func Test_addLiveSession(t *testing.T) {
  s := Session{}
  addLiveSession(&s)
  if !live_sessions[&s] {
    t.Error("Did not create live_sessions")
  }
}

func Test_removeLiveSessions(t *testing.T) {
  s := Session{}
  live_sessions[&s] = true
  removeLiveSession(&s)
  if live_sessions[&s] {
    t.Errorf("%i\n", live_sessions[&s])
  }
}

func Test_Session_Start(t *testing.T) {
  t.Skip("Session.Start not tested yet")
}

func Test_Session_listen(t *testing.T) {
  t.Skip("Session.listen not tested yet")
}

func Test_Session_Send(t *testing.T) {
  t.Skip("Session.Send not tested yet")
}

func Test_SendAll(t *testing.T) {
  t.Skip("SendAll not tested yet")
}

func Test_SendAllExcept(t *testing.T) {
  t.Skip("SendAllExcept not tested yet")
}

