# SAGO (Socket Action Go)

Sago is a framework for web development with the use of websockets. Gives the
use of simple actions to pass between server and client for real-time
communication.

## Install

```sh
$ go get github.com/craigjackson/sago
```

## Usage

Will need to import sago.
```go
import (
  "github.com/craigjackson/sago"
)
```

Need to add the actions at the first part of main() then use Run() to start the
server:
```go
func main() {
  sago.AddAction("ping", func(*sago.Request) {
    r.Session.Send(r.Id, "pong", r.Args)
  })
  sago.Run()
}
```

On "ping" from the client, we will send a reply (passing the Id in the first
parameter) JSON message back to the client with the action "pong". The extra
Args passed back will just be the exact data the client passed to us.

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

## Settings

Two methods for overriding default settings.

### 1.

For server settings, you may use in environment variables.

```sh
SERVER_HOST=example.com SERVER_PORT=80 go run app.go
```

### 2.

For all settings, you may import settings and set the variable yourself. Make
sure to establish as early in the process as possible (especially before
sago.Run()).

```go
import "github.com/craigjackson/sago/settings"

//...

func main() {
  settings.SERVER_HOST = "localhost"
  settings.SERVER_PORT = "4001"
  settings.WEBSOCKET_PATH = "/ws_now_here"
  //...
  sago.Run()
}
```

## License

The MIT License - Copyright (c) 2013 Craig Jackson

