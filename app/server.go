package main

import (

	"github.com/rs/cors"
	"github.com/zenazn/goji"
	"github.com/sinelga/images_provider/handlers"
	"log"
	"log/syslog"
//	"net/http"
)

func main() {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})
	goji.Use(c.Handler)

	goji.Get("/api", handlers.MhandleAll)
	goji.Get("/api/:id", handlers.MhandleAll)
	goji.Get("/img/:id/:imgfile/:width/:height", handlers.ImageShow)
	goji.Get("/chat/:uuid/:phone/:ask", handlers.GetChatAnswer)

	goji.Get("/fullimage/:id/original/:imgfile", handlers.ImageFullShow)
	goji.Get("/*",handlers.Elaborate)

	goji.Serve()
}
