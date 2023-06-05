package main

import (
    static "github.com/gin-contrib/static"
    "github.com/gin-gonic/gin"
    "github.com/olahol/melody"
)

func main() {
    router := gin.Default()
    m := melody.New()

    router.Use(static.Serve("/", static.LocalFile("./public", true)))

    router.GET("/websocket", func(c *gin.Context){
        m.HandleRequest(c.Writer, c.Request)
    })

    m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})

    router.Run(":5000")
}
