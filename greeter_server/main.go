package main

import(
	"context"
	"log"
	"net"

	"go-grpc/api/proto/v1"
	pb "go-grpc/api/proto/v1/"
)

const (
	port = ":50051"
)

//server is used to imolement helloworld.GreeterServer.
type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pp.HelloReply, error){
	log.printf("Receieved: %c", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main(){
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

