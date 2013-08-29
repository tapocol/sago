package settings

import (
  "os"
)

var (
  SERVER_HOST = ""
  SERVER_PORT = "4000"
  ROOT_TEMPLATE = "index.html"
  ROOT_PATH = "/"
  PUBLIC_DIR = "public"
  WEBSOCKET_PATH = "/ws"
)

func init() {
  if os.Getenv("SERVER_HOST") != "" {
    SERVER_HOST = os.Getenv("SERVER_HOST")
  }
  if os.Getenv("SERVER_PORT") != "" {
    SERVER_PORT = os.Getenv("SERVER_PORT")
  }
}

