package sago

import (
  "github.com/craigjackson/sago/settings"

  "fmt"
  "net/http"
  "text/template"
  "code.google.com/p/go.net/websocket"
)

func wsHandler(ws *websocket.Conn) {
  s := Session{ ws: ws, Data: make(map[string]interface{}) }
  s.Start()
}

//type httpHandler struct {}
//
//func (h *httpHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
//  fmt.Printf("Request: [%s] %s %s\n", request.RemoteAddr, request.Method, request.RequestURI)
//  http.DefaultServeMux.ServeHTTP(response, request)
//}

func Run() {
  indexTemplate := template.Must(template.ParseFiles(settings.ROOT_TEMPLATE))
  http.HandleFunc(settings.ROOT_PATH, func(response http.ResponseWriter, request *http.Request) {
    indexTemplate.Execute(response, request.Host)
  })
  http.Handle(settings.WEBSOCKET_PATH, websocket.Handler(wsHandler))
  fmt.Println(fmt.Sprintf("Starting Server on %s:%s ...", settings.SERVER_HOST, settings.SERVER_PORT))
  //err := http.ListenAndServe(fmt.Sprintf("%s:%s", settings.SERVER_HOST, settings.SERVER_PORT), httpHandler{})
  err := http.ListenAndServe(fmt.Sprintf("%s:%s", settings.SERVER_HOST, settings.SERVER_PORT), nil)
  if err != nil {
    panic("ListenAndServe: " + err.Error())
  }
}

