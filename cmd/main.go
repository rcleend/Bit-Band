package main

import (
    "fmt"
    "log"
    "net/http"
    "golang.org/x/net/websocket"
)

type player struct {
    id int
    x int
    y int
}

func echo(ws *websocket.Conn) {
    var err error
    player1:= player{1, 50, 50}

    for {
        var reply string

        if err = websocket.Message.Receive(ws, &reply); err != nil {
            fmt.Println("Can't receive")
            break
        }

        fmt.Println("Received back from client: " + reply)
        if reply == "hoi" {
            reply = "doei"
        }

        msg := "Received:  " + reply
        fmt.Println("Sending to client: " + msg)

        if err = websocket.Message.Send(ws, msg); err != nil {
             fmt.Println("Can't send")
             break
        }
    }
}

func main() {
    http.Handle("/", websocket.Handler(echo))
    if err := http.ListenAndServe(":80", nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}




