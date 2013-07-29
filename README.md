# SAGO (Socket Action Go)

Sago is a framework for web development with the use of websockets. Gives the
use of simple actions to pass between server and client for real-time
communication.

## Install

```sh
$ go get github.com/craigjackson/sago
```

## Usage

Will need to import sago and sago/session.
```go
import (
  "github.com/craigjackson/sago"
  "github.com/craigjackson/sago/session"
)
```

Will need to create a function that accepts three specific parameters:
```go
func ping(s *session.Session, id string, data map[string]interface{}) {
  // handle action
}
```

Need to add the actions at the first part of main() then use Run() to start the
server:
```go
func main() {
  session.AddAction("ping", ping)
  session.Run()
}
```

As you noticed, we just commented out the place to handle the action. Lets just
respond to the client with "pong":
```go
func ping(s *session.Session, id string, data map[string]interface{}) {
  s.Send(id, "pong", data)
}
```

That will send a reply (passing the id in the first parameter) JSON message back
to the client with the action "pong". The extra data passed back will just be
the exact data the client passed to us.

Now, you just need to start the server from the command-line:
```sh
$ go run app.go
```

On the client-side now, you can just use the following in javascript to see sago
in action:
```javascript
var ws = new WebSocket("ws://localhost:4000/ws");
ws.onmessage = function(message) { console.log(message) };
ws.send(JSON.stringify({ id: "sago123", action: "ping", data: { foo: "bar" } }));
```

## License

The MIT License - Copyright (c) 2013 Craig Jackson

