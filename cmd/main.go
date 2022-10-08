package main

import (
	"log"
	"sync"

	"github.com/lavish-gambhir/article-service/internal/client"
	"github.com/lavish-gambhir/article-service/internal/server"
)

func main() {
	srv := server.NewArticleApiServer()
	cl := client.NewXClient()
	p := "9090"
	var wg sync.WaitGroup
	wg.Add(1)
	log.Println("[server]::init")
	go func() {
		srv.RunServer(p)
	}()
	cl.RunClient(p)

	wg.Wait()
}
