package nexus

import (
  "fmt"
  "github.com/googollee/go-socket.io"
)
func Server() {
  server, err := socketio.NewServer(nil)

  if err != nil {
    log.Fatal(err)
  }

  server.On("connection", func(socket socketio.Socket) {
    log.Println("on connection")

    socket.On("disconnection", func() {
      log.Println("on disconnection")
    })
  })
}
