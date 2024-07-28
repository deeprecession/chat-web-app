package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/deeprecession/chat-web-app/api"
	"github.com/deeprecession/chat-web-app/api/db"
)

func main() {
	mongoURI := "mongodb://root:pass@localhost:27017"
	storage, err := db.NewMongoStorage(mongoURI)
	if err != nil {
		log.Fatalf("failed to create db connection: %s", err)
	}

	log.Printf("Connected to a MongoDB! url=%q", mongoURI)

	r := gin.Default()
	r.Use(cors.Default())

	server, err := api.NewServer(storage)
	if err != nil {
		log.Fatalln(err)
	}

	api.RegisterHandlers(r, server)

	s := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8000",
	}

	log.Fatal(s.ListenAndServe())
}
