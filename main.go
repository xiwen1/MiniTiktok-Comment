package main

import (
	service "MiniTiktok-Comment/comment"
	"MiniTiktok-Comment/proto/comment"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	grpcPort = ":15432"
)

func main() {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	err = service.InitComment()
	if err != nil {
		log.Fatalln("failed to init comment", err)
	}
	s := grpc.NewServer()
	comment.RegisterCommentActionServer(s, service.CommentActionServer{})
	log.Println("Serving gRPC on 0.0.0.0:8080")
	log.Println(s.Serve(lis))

}
