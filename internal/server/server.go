package server

import (
	"context"
	"log"
	"net"

	v1 "github.com/lavish-gambhir/article-service/pkg/v1"
	"google.golang.org/grpc"
)

type ArticleApiServer struct {
	v1.UnsafeArticleServiceServer
}

func NewArticleApiServer() *ArticleApiServer { return &ArticleApiServer{} }

func (a *ArticleApiServer) GetArticles(ctx context.Context, rq *v1.GetArticleRequest) (*v1.ArticleResponse, error) {
	return &v1.ArticleResponse{
		Id:     "12101",
		Title:  "Fall of social media giants",
		Name:   "Fb",
		Author: "Cal Nweport",
	}, nil
}

func (a *ArticleApiServer) RunServer(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("unable to create listener: %v", err)
	}

	server := grpc.NewServer()
	v1.RegisterArticleServiceServer(server, NewArticleApiServer())
	log.Println("starting Article gRPC service..")
	server.Serve(listener)
}
