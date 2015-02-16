package nexus

import (
	"log"
  "fmt"
	"net/http"

	"github.com/googollee/go-socket.io"
)

func StartServe(port int) {
	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", func(socket socketio.Socket) {
		log.Println("on connection")

    m := make(map[string]interface{})
    m["a"] = "你好"
    socket.Emit("hello", m)

		socket.On("disconnection", func() {
			log.Println("on disconnection")
		})
	})

  server.On("error", func(socket socketio.Socket, err error) {
    log.Println("error: ", err)
  })

  http.Handle("/", server)
  log.Printf("Serving at http://localhost:%d/daemon/", port)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
