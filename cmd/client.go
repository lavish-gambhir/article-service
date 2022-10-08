package main

import (
	"context"
	"log"

	v1 "github.com/lavish-gambhir/article-service/pkg/v1"
	"google.golang.org/grpc"
)

func RunClient(port string) {
	conn, err := grpc.Dial(":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	defer conn.Close()
	cl := v1.NewArticleServiceClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	resp, err := cl.GetArticles(ctx, &v1.GetArticleRequest{UserId: 1121})
	if err != nil {
		log.Fatalf("unable to get articles: %v", err)
	}
	log.Println("******got article response********")
	log.Printf("id = %s, title = %s", resp.Id, resp.Title)
}
