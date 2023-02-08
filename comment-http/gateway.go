package main

import (
	"MiniTiktok-Comment/proto/comment"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

var gwPort = ":54321"

func main() {
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:15432",
		//grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server", err)
	}
	connComment := comment.NewCommentActionClient(conn)

	gwmux := runtime.NewServeMux()

	err = comment.RegisterCommentActionHandlerClient(context.Background(), gwmux, connComment)
	if err != nil {
		log.Fatalln("failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    gwPort,
		Handler: gwmux,
	}

	log.Printf("serving grpc-gateway on http://0.0.0.0:%v", gwPort)
	log.Println(gwServer.ListenAndServe())

}
