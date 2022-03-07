package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcBlog/blog/blog_pb"
	"grpcBlog/blog/blog_server"
	"grpcBlog/blog/config"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {
	// If program crashes we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Blog Service started")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	var opts []grpc.ServerOption

	s := grpc.NewServer(opts...)
	blog_pb.RegisterBlogServiceServer(s, &blog_server.Server{})
	go func() {
		fmt.Println("Starting Server ....")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v\n", err)
		}
	}()
	fmt.Println("Server started ...")

	// Connect to MongoDB
	config.ConnectDB()

	// Wait for ctrl c to exit
	ch := make(chan os.Signal, 1)
	// pass os signal to our channel
	signal.Notify(ch, os.Interrupt)

	//block until channel receives signal
	<-ch
	fmt.Println("Stopping Server ....")
	s.GracefulStop()
	fmt.Println("Closing Listener ...")
	listenerErr := lis.Close()
	if listenerErr != nil {
		return
	}
	fmt.Println("Disconnecting mongodb")
	err = config.DATABASE.Disconnect(context.TODO())
	if err != nil {
		return
	}

}
